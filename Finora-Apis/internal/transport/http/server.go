package httptransport

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/status"

	finorav1 "finoraai/backend/api/proto/finora/v1"
	"finoraai/backend/internal/config"
)

// Service is the subset of finora.v1.FinoraServiceServer used by the HTTP gateway.
type Service interface {
	Login(ctx context.Context, req *finorav1.LoginRequest) (*finorav1.LoginResponse, error)
	Health(ctx context.Context, req *finorav1.HealthRequest) (*finorav1.HealthResponse, error)
}

// NewServer builds and returns an *http.Server ready to Serve.
func NewServer(cfg config.Config, log *zap.Logger, svc Service) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/auth/login", loginHandler(log, svc))
	mux.HandleFunc("/api/health", healthHandler(log, svc))

	return &http.Server{
		Addr:         net.JoinHostPort("0.0.0.0", cfg.HTTPPort),
		Handler:      cors(cfg)(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

// cors adds CORS headers and handles preflight OPTIONS requests.
func cors(cfg config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.Header().Set("Access-Control-Max-Age", "86400")
				if !cfg.IsProduction() {
					w.Header().Set("Access-Control-Allow-Credentials", "true")
				}
			}
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func loginHandler(log *zap.Logger, svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}
		resp, err := svc.Login(r.Context(), &finorav1.LoginRequest{
			Email:    strings.TrimSpace(body.Email),
			Password: body.Password,
		})
		if err != nil {
			code, msg := grpcErrToHTTP(err)
			writeJSON(w, code, map[string]string{"error": msg})
			log.Warn("login failed", zap.String("email", body.Email), zap.Error(err))
			return
		}
		writeJSON(w, http.StatusOK, resp)
	}
}

func healthHandler(log *zap.Logger, svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := svc.Health(r.Context(), &finorav1.HealthRequest{})
		if err != nil {
			log.Error("health check failed", zap.Error(err))
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "unhealthy"})
			return
		}
		writeJSON(w, http.StatusOK, resp)
	}
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

// grpcErrToHTTP maps gRPC status codes to HTTP status codes.
func grpcErrToHTTP(err error) (int, string) {
	st, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, "internal error"
	}
	switch st.Code() {
	case 3: // InvalidArgument
		return http.StatusBadRequest, st.Message()
	case 5: // NotFound
		return http.StatusNotFound, st.Message()
	case 7: // PermissionDenied
		return http.StatusForbidden, st.Message()
	case 9: // FailedPrecondition
		return http.StatusServiceUnavailable, st.Message()
	case 16: // Unauthenticated
		return http.StatusUnauthorized, st.Message()
	default:
		return http.StatusInternalServerError, "internal error"
	}
}
