package auth

import "context"

type claimsCtxKey struct{}

// WithClaims attaches verified JWT claims to the context for downstream handlers.
func WithClaims(ctx context.Context, c *Claims) context.Context {
	return context.WithValue(ctx, claimsCtxKey{}, c)
}

// ClaimsFromContext returns claims set by the JWT interceptor.
func ClaimsFromContext(ctx context.Context) (*Claims, bool) {
	v, ok := ctx.Value(claimsCtxKey{}).(*Claims)
	return v, ok
}
