package grpctransport

import (
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	finorav1 "finoraai/backend/api/proto/finora/v1"
	"finoraai/backend/internal/config"
	"finoraai/backend/internal/transport/grpc/interceptor"
	"finoraai/backend/internal/transport/grpc/routing"
)

// Deps are gRPC service implementations registered on the server.
type Deps struct {
	Finora    finorav1.FinoraServiceServer
	Health    *health.Server
	AuthUnary grpc.UnaryServerInterceptor // optional JWT (or other) auth
}

// NewServer constructs a *grpc.Server with shared interceptors and standard registrations.
func NewServer(cfg config.Config, log *zap.Logger, deps Deps) (*grpc.Server, error) {
	if deps.Finora == nil {
		return nil, fmt.Errorf("deps.Finora is required")
	}
	if deps.Health == nil {
		return nil, fmt.Errorf("deps.Health is required")
	}

	chain := []grpc.UnaryServerInterceptor{
		interceptor.Recovery(log),
		interceptor.RequestID(),
	}
	if deps.AuthUnary != nil {
		chain = append(chain, deps.AuthUnary)
	}
	chain = append(chain, interceptor.UnaryLogging(log))

	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(cfg.MaxRecvBytes()),
		grpc.MaxSendMsgSize(cfg.MaxSendBytes()),
		grpc.ChainUnaryInterceptor(chain...),
	}

	s := grpc.NewServer(opts...)

	finorav1.RegisterFinoraServiceServer(s, deps.Finora)
	grpc_health_v1.RegisterHealthServer(s, deps.Health)

	deps.Health.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	deps.Health.SetServingStatus(routing.FinoraServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	if cfg.EnableReflection {
		reflection.Register(s)
	}

	return s, nil
}
