package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Okemwag/giftbox/internal/platform/config"
	"github.com/pressly/goose/v3"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const migrationsDir = "db/migrations"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(2)
	}

	cfg := config.LoadConfig()
	if cfg.DatabaseDSN == "" {
		fmt.Fprintln(os.Stderr, "DATABASE_DSN is required to run migrations")
		os.Exit(1)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		fmt.Fprintf(os.Stderr, "set goose dialect: %v\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "connect database: %v\n", err)
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]
	if err := run(command, db, args); err != nil {
		fmt.Fprintf(os.Stderr, "migration %s failed: %v\n", command, err)
		os.Exit(1)
	}
}

func run(command string, db *sql.DB, args []string) error {
	switch command {
	case "up":
		return goose.Up(db, migrationsDir)
	case "up-by-one":
		return goose.UpByOne(db, migrationsDir)
	case "down":
		return goose.Down(db, migrationsDir)
	case "down-to":
		if len(args) != 1 {
			return fmt.Errorf("down-to requires a target version")
		}
		version, err := parseVersion(args[0])
		if err != nil {
			return err
		}
		return goose.DownTo(db, migrationsDir, version)
	case "redo":
		return goose.Redo(db, migrationsDir)
	case "reset":
		return goose.Reset(db, migrationsDir)
	case "status":
		return goose.Status(db, migrationsDir)
	case "version":
		return goose.Version(db, migrationsDir)
	default:
		return fmt.Errorf("unknown command %q", command)
	}
}

func parseVersion(value string) (int64, error) {
	var version int64
	if _, err := fmt.Sscanf(value, "%d", &version); err != nil {
		return 0, fmt.Errorf("invalid version %q", value)
	}
	return version, nil
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "usage: go run ./cmd/migrate <command>")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "commands:")
	fmt.Fprintln(os.Stderr, "  up        apply all pending migrations")
	fmt.Fprintln(os.Stderr, "  up-by-one apply the next pending migration")
	fmt.Fprintln(os.Stderr, "  down      roll back the latest migration")
	fmt.Fprintln(os.Stderr, "  down-to   roll back to a specific version")
	fmt.Fprintln(os.Stderr, "  redo      roll back and re-apply the latest migration")
	fmt.Fprintln(os.Stderr, "  reset     roll back all migrations")
	fmt.Fprintln(os.Stderr, "  status    show migration status")
	fmt.Fprintln(os.Stderr, "  version   show current migration version")
}
