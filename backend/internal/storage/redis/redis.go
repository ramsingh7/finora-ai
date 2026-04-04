package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Connect dials Redis and verifies with PING.
func Connect(ctx context.Context, addr, password string, db int) (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     32,
		MinIdleConns: 4,
	})
	if err := c.Ping(ctx).Err(); err != nil {
		_ = c.Close()
		return nil, fmt.Errorf("ping redis: %w", err)
	}
	return c, nil
}
