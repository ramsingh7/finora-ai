package interceptor

import (
	"context"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"finoraai/backend/internal/auth"
	"finoraai/backend/internal/constant"
	"finoraai/backend/internal/transport/grpc/routing"
)

// JWTAuth enforces Bearer JWT on all RPCs except health, reflection, and Login.
func JWTAuth(log *zap.Logger, parser *auth.JWT) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if routing.IsPublicMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}
		raw := ""
		if v := md.Get(constant.MetadataAuthorization); len(v) > 0 {
			raw = v[0]
		}
		token := extractBearer(raw)
		if token == "" {
			return nil, status.Error(codes.Unauthenticated, "missing bearer token")
		}
		claims, err := parser.Parse(token)
		if err != nil {
			log.Debug("jwt parse failed", zap.Error(err))
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}
		ctx = auth.WithClaims(ctx, claims)
		return handler(ctx, req)
	}
}

func extractBearer(h string) string {
	h = strings.TrimSpace(h)
	if h == "" {
		return ""
	}
	parts := strings.SplitN(h, " ", 2)
	if len(parts) != 2 {
		return ""
	}
	if !strings.EqualFold(parts[0], "bearer") {
		return ""
	}
	return strings.TrimSpace(parts[1])
}
