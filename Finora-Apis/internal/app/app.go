package app

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"

	"finoraai/backend/internal/auth"
	"finoraai/backend/internal/config"
	"finoraai/backend/internal/filemanager"
	finorasvc "finoraai/backend/internal/service/finora"
	"finoraai/backend/internal/storage/postgres"
	redisstore "finoraai/backend/internal/storage/redis"
	grpctransport "finoraai/backend/internal/transport/grpc"
	"finoraai/backend/internal/transport/grpc/interceptor"
	"finoraai/backend/internal/user"
	"finoraai/backend/internal/version"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

// Run boots dependencies, serves gRPC until ctx is cancelled, then shuts down gracefully.
func Run(ctx context.Context, cfg config.Config, log *zap.Logger) error {
	fm, err := filemanager.New(cfg.DataDir, cfg.MaxUploadBytes())
	if err != nil {
		return fmt.Errorf("file manager: %w", err)
	}

	var pool *pgxpool.Pool
	if cfg.DatabaseURL != "" {
		pool, err = postgres.Connect(ctx, cfg.DatabaseURL)
		if err != nil {
			return fmt.Errorf("postgres: %w", err)
		}
		defer pool.Close()
		log.Info("connected to postgres")
	} else {
		log.Warn("DATABASE_URL empty; Login and user features disabled")
	}

	var rdb *redis.Client
	if cfg.RedisAddr != "" {
		rdb, err = redisstore.Connect(ctx, cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
		if err != nil {
			return fmt.Errorf("redis: %w", err)
		}
		defer func() { _ = rdb.Close() }()
		log.Info("connected to redis", zap.String("addr", cfg.RedisAddr))
	} else {
		log.Warn("REDIS_ADDR empty; session cache hooks disabled")
	}

	var userStore user.Store
	if pool != nil {
		userStore = user.NewPostgresStore(pool)
	}

	var jwtSvc *auth.JWT
	if cfg.JWTSecret != "" {
		jwtSvc, err = auth.NewJWT(cfg.JWTSecret, cfg.JWTIssuer, cfg.JWTTTL())
		if err != nil {
			return fmt.Errorf("jwt: %w", err)
		}
	} else {
		log.Warn("JWT_SECRET empty; Login will fail and JWT interceptor disabled")
	}

	var authUnary grpc.UnaryServerInterceptor
	if jwtSvc != nil {
		authUnary = interceptor.JWTAuth(log, jwtSvc)
	}

	healthSrv := health.NewServer()
	finoraSvc := finorasvc.NewService(finorasvc.Params{
		Log:     log,
		Version: version.String(),
		Files:   fm,
		Users:   userStore,
		JWT:     jwtSvc,
		Redis:   rdb,
	})

	srv, err := grpctransport.NewServer(cfg, log, grpctransport.Deps{
		Finora:    finoraSvc,
		Health:    healthSrv,
		AuthUnary: authUnary,
	})
	if err != nil {
		return err
	}

	addr := net.JoinHostPort(cfg.GRPCHost, cfg.GRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen %s: %w", addr, err)
	}

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Error("grpc Serve returned", zap.Error(err))
		}
	}()

	log.Info("gRPC listening",
		zap.String("addr", lis.Addr().String()),
		zap.String("version", version.Full()),
		zap.String("data_dir", fm.Root()),
	)

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownGrace())
	defer cancel()

	done := make(chan struct{})
	go func() {
		srv.GracefulStop()
		close(done)
	}()

	select {
	case <-shutdownCtx.Done():
		log.Warn("graceful stop timed out, forcing Stop")
		srv.Stop()
	case <-done:
	}

	log.Info("gRPC server stopped")
	return nil
}
