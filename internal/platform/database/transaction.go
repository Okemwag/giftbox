package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Transactor interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context, tx pgx.Tx) error) error
}

func (p *Postgres) WithinTransaction(ctx context.Context, fn func(ctx context.Context, tx pgx.Tx) error) error {
	tx, err := p.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	if err := fn(ctx, tx); err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return fmt.Errorf("transaction failed: %w; rollback failed: %w", err, rollbackErr)
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil
}
