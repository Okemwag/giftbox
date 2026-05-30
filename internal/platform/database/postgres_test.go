package database

import (
	"context"
	"strings"
	"testing"
)

func TestOpenRequiresDatabaseURL(t *testing.T) {
	_, err := Open(context.Background(), Config{})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "database url is required") {
		t.Fatalf("expected missing database URL error, got %v", err)
	}
}

func TestOpenRejectsInvalidDatabaseURL(t *testing.T) {
	_, err := Open(context.Background(), Config{DatabaseURL: "://bad"})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "parse database url") {
		t.Fatalf("expected parse database URL error, got %v", err)
	}
}
