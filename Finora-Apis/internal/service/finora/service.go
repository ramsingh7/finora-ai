package finora

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	finorav1 "finoraai/backend/api/proto/finora/v1"
	"finoraai/backend/internal/auth"
	"finoraai/backend/internal/constant"
	"finoraai/backend/internal/filemanager"
	"finoraai/backend/internal/user"

	"github.com/redis/go-redis/v9"
)

// Params wires dependencies for FinoraService.
type Params struct {
	Log     *zap.Logger
	Version string
	Files   *filemanager.Manager
	Users   user.Store
	JWT     *auth.JWT
	Redis   *redis.Client
}

// Service implements finora.v1.FinoraService.
type Service struct {
	finorav1.UnimplementedFinoraServiceServer
	log     *zap.Logger
	version string
	files   *filemanager.Manager
	users   user.Store
	jwt     *auth.JWT
	redis   *redis.Client
}

// NewService constructs the service from Params.
func NewService(p Params) *Service {
	return &Service{
		log:     p.Log,
		version: p.Version,
		files:   p.Files,
		users:   p.Users,
		jwt:     p.JWT,
		redis:   p.Redis,
	}
}

// Health implements the public health check used by gateways and clients.
func (s *Service) Health(ctx context.Context, _ *finorav1.HealthRequest) (*finorav1.HealthResponse, error) {
	_ = ctx
	return &finorav1.HealthResponse{
		Status:  "ok",
		Version: s.version,
	}, nil
}

// Login exchanges email/password for a JWT access token.
func (s *Service) Login(ctx context.Context, req *finorav1.LoginRequest) (*finorav1.LoginResponse, error) {
	if s.users == nil {
		return nil, status.Error(codes.FailedPrecondition, "user store not configured")
	}
	if s.jwt == nil {
		return nil, status.Error(codes.FailedPrecondition, "JWT not configured; set JWT_SECRET")
	}

	email := strings.TrimSpace(strings.ToLower(req.GetEmail()))
	if email == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password required")
	}

	u, err := s.users.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return nil, status.Error(codes.Unauthenticated, "invalid credentials")
		}
		s.log.Error("login lookup failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "login failed")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.GetPassword())); err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	token, exp, err := s.jwt.Sign(u.ID, u.Email)
	if err != nil {
		s.log.Error("jwt sign failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "login failed")
	}

	if s.redis != nil {
		ttl := time.Until(exp)
		if ttl > 0 {
			_ = s.redis.Set(ctx, "finora:login:"+u.ID, time.Now().UTC().Format(time.RFC3339Nano), ttl).Err()
		}
	}

	sec := int64(time.Until(exp).Seconds())
	if sec < 1 {
		sec = 1
	}
	return &finorav1.LoginResponse{
		AccessToken:      token,
		ExpiresInSeconds: sec,
	}, nil
}

// Files exposes the path-safe file manager for other internal packages.
func (s *Service) Files() *filemanager.Manager { return s.files }

// Log returns the service logger.
func (s *Service) Log() *zap.Logger { return s.log }

// AppName returns the constant application name for payloads.
func (s *Service) AppName() string { return constant.AppName }
