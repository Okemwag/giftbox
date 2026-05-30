package errors

import (
	"net/http"
	"testing"
)

func TestToHTTPMapsApplicationErrors(t *testing.T) {
	status, payload := ToHTTP(New(CodeConsentRequired, "marketing consent is required"))
	if status != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d", status)
	}
	if payload.Code != CodeConsentRequired {
		t.Fatalf("expected code %s, got %s", CodeConsentRequired, payload.Code)
	}
}

func TestToHTTPHidesUnknownErrors(t *testing.T) {
	status, payload := ToHTTP(assertionError("database connection failed"))
	if status != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", status)
	}
	if payload.Message != "internal server error" {
		t.Fatalf("expected safe message, got %q", payload.Message)
	}
}

type assertionError string

func (e assertionError) Error() string {
	return string(e)
}
