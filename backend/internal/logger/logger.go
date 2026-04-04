package logger

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New builds a zap.Logger for the given level and environment name.
func New(level, env string) (*zap.Logger, error) {
	lvl, err := parseLevel(level)
	if err != nil {
		return nil, err
	}

	if strings.EqualFold(env, "development") || strings.EqualFold(env, "dev") || strings.EqualFold(env, "local") {
		cfg := zap.NewDevelopmentConfig()
		cfg.Level = zap.NewAtomicLevelAt(lvl)
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return cfg.Build()
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(lvl)
	return cfg.Build()
}

func parseLevel(s string) (zapcore.Level, error) {
	var l zapcore.Level
	if err := l.UnmarshalText([]byte(strings.ToLower(s))); err != nil {
		return l, fmt.Errorf("log level: %w", err)
	}
	return l, nil
}
