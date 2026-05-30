package main

import (
	"fmt"
	"os"

	"github.com/Okemwag/giftbox/internal/platform"
)

func main() {
	cfg := platform.LoadConfig()
	if cfg.DatabaseDSN == "" {
		fmt.Fprintln(os.Stderr, "DATABASE_DSN is required to run migrations")
		os.Exit(1)
	}

	fmt.Printf("migration runner scaffold ready for %s\n", cfg.DatabaseDSN)
}
