package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"finoraai/backend/internal/app"
	"finoraai/backend/internal/config"
	"finoraai/backend/internal/logger"
)

func main() {
	// Optional local overrides; missing .env is fine in containers / production.
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config: %v\n", err)
		os.Exit(1)
	}

	log, err := logger.New(cfg.LogLevel, cfg.Env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "logger: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = log.Sync() }()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx, cfg, log); err != nil {
		log.Fatal("application error", zap.Error(err))
	}
}
