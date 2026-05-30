# Database Migrations

Giftbox uses [Goose](https://pressly.github.io/goose/) for database migrations.

## Why Goose

Goose keeps migration changes explicit and reviewable. That matters for this platform because tenant isolation, idempotent payment ingestion, loyalty ledgers, audit logs, and outbox events all depend on exact PostgreSQL constraints and indexes.

## File Format

Use one SQL file per migration:

```text
db/migrations/
└── 000001_init_core_schema.sql
```

Every SQL migration must include a `-- +goose Up` section. Reversible migrations should also include `-- +goose Down`.

```sql
-- +goose Up
CREATE TABLE example_records (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS example_records;
```

Goose annotations must be on their own line and should not be indented.

For complex SQL blocks such as PL/pgSQL functions or triggers, wrap the statement with `StatementBegin` and `StatementEnd` so Goose treats the block as one statement:

```sql
-- +goose Up
-- +goose StatementBegin
CREATE FUNCTION example_trigger()
RETURNS trigger AS $$
BEGIN
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
```

## Running Migrations

Set `DATABASE_DSN`, then run the migration command:

```sh
DATABASE_DSN=postgres://giftbox:giftbox@localhost:5432/giftbox?sslmode=disable go run ./cmd/migrate status
DATABASE_DSN=postgres://giftbox:giftbox@localhost:5432/giftbox?sslmode=disable go run ./cmd/migrate up
DATABASE_DSN=postgres://giftbox:giftbox@localhost:5432/giftbox?sslmode=disable go run ./cmd/migrate down
```

Supported commands:

```text
up
up-by-one
down
down-to <version>
redo
reset
status
version
```

## Operational Rules

- Keep schema migrations small and reviewable.
- Prefer SQL migrations over Go migrations for schema changes.
- Add `tenant_id` to every tenant-owned table.
- Add uniqueness constraints for idempotency at the database layer.
- Never edit a migration that has been applied outside local development; add a new migration instead.
- Avoid irreversible `Down` sections unless the migration truly cannot be reversed safely.
- Use transactions for migrations unless a PostgreSQL operation explicitly requires `NO TRANSACTION`.
