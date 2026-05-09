package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/caarlos0/env/v11"
)

// Config is loaded from environment variables (12-factor); optional .env in dev via godotenv in main.
type Config struct {
	Env              string `env:"ENV" envDefault:"development"`
	GRPCHost         string `env:"GRPC_HOST" envDefault:"0.0.0.0"`
	GRPCPort         string `env:"GRPC_PORT" envDefault:"50051"`
	HTTPPort         string `env:"HTTP_PORT" envDefault:"3000"`
	LogLevel         string `env:"LOG_LEVEL" envDefault:"info"`
	DataDir          string `env:"DATA_DIR" envDefault:"./data"`
	MaxUploadMB      int    `env:"MAX_UPLOAD_MB" envDefault:"32"`
	ShutdownGraceSec int    `env:"SHUTDOWN_GRACE_SECONDS" envDefault:"15"`
	EnableReflection bool   `env:"GRPC_REFLECTION" envDefault:"true"`
	MaxRecvMB        int    `env:"GRPC_MAX_RECV_MB" envDefault:"4"`
	MaxSendMB        int    `env:"GRPC_MAX_SEND_MB" envDefault:"4"`
	DatabaseURL      string `env:"DATABASE_URL" envDefault:""`
	RedisAddr        string `env:"REDIS_ADDR" envDefault:""`
	RedisPassword    string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDB          int    `env:"REDIS_DB" envDefault:"0"`
	JWTSecret        string `env:"JWT_SECRET" envDefault:""`
	JWTIssuer        string `env:"JWT_ISSUER" envDefault:"finora-backend"`
	JWTExpireMinutes int    `env:"JWT_EXPIRE_MINUTES" envDefault:"60"`
}

// Load parses environment into Config and validates derived fields.
func Load() (Config, error) {
	var c Config
	if err := env.Parse(&c); err != nil {
		return Config{}, fmt.Errorf("parse env: %w", err)
	}
	if err := c.Validate(); err != nil {
		return Config{}, err
	}
	return c, nil
}

// Validate enforces invariants after env parsing.
func (c Config) Validate() error {
	if c.MaxUploadMB < 1 || c.MaxUploadMB > 1024 {
		return fmt.Errorf("MAX_UPLOAD_MB must be between 1 and 1024, got %d", c.MaxUploadMB)
	}
	if c.ShutdownGraceSec < 1 || c.ShutdownGraceSec > 300 {
		return fmt.Errorf("SHUTDOWN_GRACE_SECONDS must be between 1 and 300, got %d", c.ShutdownGraceSec)
	}
	if _, err := strconv.Atoi(c.GRPCPort); err != nil || c.GRPCPort == "" {
		return fmt.Errorf("GRPC_PORT must be a numeric port, got %q", c.GRPCPort)
	}
	if c.MaxRecvMB < 1 || c.MaxRecvMB > 256 {
		return fmt.Errorf("GRPC_MAX_RECV_MB must be between 1 and 256, got %d", c.MaxRecvMB)
	}
	if c.MaxSendMB < 1 || c.MaxSendMB > 256 {
		return fmt.Errorf("GRPC_MAX_SEND_MB must be between 1 and 256, got %d", c.MaxSendMB)
	}
	switch c.LogLevel {
	case "debug", "info", "warn", "error":
	default:
		return fmt.Errorf("LOG_LEVEL must be debug|info|warn|error, got %q", c.LogLevel)
	}
	if c.RedisDB < 0 || c.RedisDB > 15 {
		return fmt.Errorf("REDIS_DB must be 0..15, got %d", c.RedisDB)
	}
	if c.JWTExpireMinutes < 1 || c.JWTExpireMinutes > 10080 {
		return fmt.Errorf("JWT_EXPIRE_MINUTES must be between 1 and 10080, got %d", c.JWTExpireMinutes)
	}
	if c.IsProduction() {
		if c.JWTSecret == "" || len(c.JWTSecret) < 32 {
			return fmt.Errorf("production requires JWT_SECRET of at least 32 characters")
		}
		if c.DatabaseURL == "" {
			return fmt.Errorf("production requires DATABASE_URL")
		}
	}
	return nil
}

// MaxUploadBytes returns MaxUploadMB as bytes for storage quotas.
func (c Config) MaxUploadBytes() int64 {
	return int64(c.MaxUploadMB) * 1024 * 1024
}

// MaxRecvBytes returns the configured max receive message size in bytes.
func (c Config) MaxRecvBytes() int {
	return c.MaxRecvMB * 1024 * 1024
}

// MaxSendBytes returns the configured max send message size in bytes.
func (c Config) MaxSendBytes() int {
	return c.MaxSendMB * 1024 * 1024
}

// ShutdownGrace returns shutdown timeout as duration.
func (c Config) ShutdownGrace() time.Duration {
	return time.Duration(c.ShutdownGraceSec) * time.Second
}

// IsProduction returns true when ENV is production-like.
func (c Config) IsProduction() bool {
	return c.Env == "production" || c.Env == "prod"
}

// JWTTTL returns access token lifetime.
func (c Config) JWTTTL() time.Duration {
	return time.Duration(c.JWTExpireMinutes) * time.Minute
}
