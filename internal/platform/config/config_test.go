package config

import (
	"strings"
	"testing"
)

func TestValidateRejectsMissingDatabaseURL(t *testing.T) {
	cfg := Config{
		Environment: "local",
		HTTPPort:    "8080",
		AppBaseURL:  "http://localhost:8080",
		JWTIssuer:   "http://localhost:8080",
		JWTAudience: "giftbox-api",
		LogLevel:    "info",
	}

	err := cfg.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}
	if !strings.Contains(err.Error(), "DATABASE_URL is required") {
		t.Fatalf("expected missing database URL error, got %v", err)
	}
}

func TestValidateRejectsUnsupportedEnvironment(t *testing.T) {
	cfg := validConfig()
	cfg.Environment = "qa"

	err := cfg.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}
	if !strings.Contains(err.Error(), "APP_ENV") {
		t.Fatalf("expected environment error, got %v", err)
	}
}

func TestValidateRejectsInvalidJWTIssuer(t *testing.T) {
	cfg := validConfig()
	cfg.JWTIssuer = "not-a-url"

	err := cfg.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}
	if !strings.Contains(err.Error(), "JWT_ISSUER") {
		t.Fatalf("expected JWT issuer error, got %v", err)
	}
}

func validConfig() Config {
	return Config{
		Environment: "local",
		HTTPPort:    "8080",
		DatabaseURL: "postgres://giftbox:giftbox@localhost:5432/giftbox?sslmode=disable",
		AppBaseURL:  "http://localhost:8080",
		JWTIssuer:   "http://localhost:8080",
		JWTAudience: "giftbox-api",
		LogLevel:    "info",
	}
}
