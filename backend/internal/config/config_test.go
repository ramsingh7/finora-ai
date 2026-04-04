package config

import "testing"

func TestMaxUploadBytes(t *testing.T) {
	c := Config{MaxUploadMB: 10}
	if got := c.MaxUploadBytes(); got != 10*1024*1024 {
		t.Fatalf("MaxUploadBytes: got %d", got)
	}
}

func TestValidate_OK(t *testing.T) {
	c := Config{
		MaxUploadMB:      32,
		ShutdownGraceSec: 15,
		GRPCPort:         "50051",
		LogLevel:         "info",
		MaxRecvMB:        4,
		MaxSendMB:        4,
		RedisDB:          0,
		JWTExpireMinutes: 60,
	}
	if err := c.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestValidate_BadLogLevel(t *testing.T) {
	c := Config{
		MaxUploadMB:      32,
		ShutdownGraceSec: 15,
		GRPCPort:         "50051",
		LogLevel:         "trace",
		MaxRecvMB:        4,
		MaxSendMB:        4,
		RedisDB:          0,
		JWTExpireMinutes: 60,
	}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for invalid LOG_LEVEL")
	}
}

func TestValidate_BadPort(t *testing.T) {
	c := Config{
		MaxUploadMB:      32,
		ShutdownGraceSec: 15,
		GRPCPort:         "not-a-port",
		LogLevel:         "info",
		MaxRecvMB:        4,
		MaxSendMB:        4,
		RedisDB:          0,
		JWTExpireMinutes: 60,
	}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for invalid GRPC_PORT")
	}
}

func TestValidate_ProductionRequiresJWT(t *testing.T) {
	c := Config{
		Env:              "production",
		DatabaseURL:      "postgres://u:p@localhost:5432/db",
		JWTSecret:        "tooshort",
		MaxUploadMB:      32,
		ShutdownGraceSec: 15,
		GRPCPort:         "50051",
		LogLevel:         "info",
		MaxRecvMB:        4,
		MaxSendMB:        4,
		RedisDB:          0,
		JWTExpireMinutes: 60,
	}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for short JWT_SECRET in production")
	}
}
