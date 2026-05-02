package interceptor

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"finoraai/backend/internal/constant"
)

// RequestID ensures every RPC has a correlation ID in context and echoes it on response metadata.
func RequestID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		id := ""
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if v := md.Get(constant.MetadataRequestID); len(v) > 0 && v[0] != "" {
				id = v[0]
			}
		}
		if id == "" {
			id = uuid.NewString()
		}
		ctx = constant.WithRequestID(ctx, id)
		_ = grpc.SetHeader(ctx, metadata.Pairs(constant.MetadataRequestID, id))
		return handler(ctx, req)
	}
}
