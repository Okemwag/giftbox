package database

import (
	"context"
	"errors"
)

type Config struct {
	DSN string
}

type Store struct {
	dsn string
}

func Open(ctx context.Context, cfg Config) (*Store, error) {
	if cfg.DSN == "" {
		return nil, errors.New("database dsn is required")
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &Store{dsn: cfg.DSN}, nil
	}
}

func (s *Store) DSN() string {
	return s.dsn
}
