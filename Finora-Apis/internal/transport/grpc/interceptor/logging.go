package interceptor

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"finoraai/backend/internal/constant"
)

// UnaryLogging logs method, duration, request ID, and outcome.
func UnaryLogging(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		fields := []zap.Field{
			zap.String("grpc.method", info.FullMethod),
			zap.Duration("duration", time.Since(start)),
		}
		if id := constant.RequestIDFromContext(ctx); id != "" {
			fields = append(fields, zap.String("request_id", id))
		}
		if err != nil {
			if s, ok := status.FromError(err); ok {
				fields = append(fields, zap.String("grpc.code", s.Code().String()))
			}
			log.Warn("grpc unary finished with error", append(fields, zap.Error(err))...)
			return resp, err
		}
		log.Info("grpc unary ok", append(fields, zap.String("grpc.code", codes.OK.String()))...)
		return resp, nil
	}
}
