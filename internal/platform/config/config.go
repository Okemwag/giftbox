package config

import (
	"fmt"
	"strings"
)

type Config struct {
	Environment string
	HTTPPort    string
	DatabaseURL string
	AppBaseURL  string
	JWTIssuer   string
	JWTAudience string
	LogLevel    string
}

func Load() (Config, error) {
	cfg := Config{
		Environment: env("APP_ENV", "local"),
		HTTPPort:    env("HTTP_PORT", "8080"),
		DatabaseURL: env("DATABASE_URL", ""),
		AppBaseURL:  env("APP_BASE_URL", "http://localhost:8080"),
		JWTIssuer:   env("JWT_ISSUER", ""),
		JWTAudience: env("JWT_AUDIENCE", ""),
		LogLevel:    env("LOG_LEVEL", "info"),
	}
	return cfg, cfg.Validate()
}

func (c Config) HTTPAddr() string {
	return ":" + strings.TrimPrefix(c.HTTPPort, ":")
}

func (c Config) Validate() error {
	var problems []string

	if !isSupportedEnvironment(c.Environment) {
		problems = append(problems, "APP_ENV must be one of local, development, staging, production, test")
	}
	if strings.TrimSpace(c.HTTPPort) == "" {
		problems = append(problems, "HTTP_PORT is required")
	}
	if strings.TrimSpace(c.DatabaseURL) == "" {
		problems = append(problems, "DATABASE_URL is required")
	}
	if err := validateURL("APP_BASE_URL", c.AppBaseURL); err != nil {
		problems = append(problems, err.Error())
	}
	if err := validateURL("JWT_ISSUER", c.JWTIssuer); err != nil {
		problems = append(problems, err.Error())
	}
	if strings.TrimSpace(c.JWTAudience) == "" {
		problems = append(problems, "JWT_AUDIENCE is required")
	}
	if !isSupportedLogLevel(c.LogLevel) {
		problems = append(problems, "LOG_LEVEL must be one of debug, info, warn, error")
	}

	if len(problems) > 0 {
		return fmt.Errorf("invalid configuration: %s", strings.Join(problems, "; "))
	}
	return nil
}
