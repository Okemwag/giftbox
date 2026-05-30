package contract

import (
	"os"
	"testing"
)

func TestOpenAPIContractExists(t *testing.T) {
	if _, err := os.Stat("../../api/openapi.yaml"); err != nil {
		t.Fatalf("expected openapi contract: %v", err)
	}
}
