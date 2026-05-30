package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	apperrors "github.com/Okemwag/giftbox/internal/shared/errors"
)

func TestDecodeJSONReturnsControlledErrorForMalformedPayload(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":`))

	var payload struct {
		Name string `json:"name"`
	}
	err := DecodeJSON(request, &payload)
	if err == nil {
		t.Fatal("expected error")
	}

	status, response := apperrors.ToHTTP(err)
	if status != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", status)
	}
	if response.Code != apperrors.CodeInvalidInput {
		t.Fatalf("expected invalid input code, got %s", response.Code)
	}
}

func TestWriteErrorReturnsConsistentJSONEnvelope(t *testing.T) {
	response := httptest.NewRecorder()

	WriteError(response, apperrors.New(apperrors.CodeNotFound, "customer not found"))

	if response.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", response.Code)
	}
	if got := response.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected JSON content type, got %q", got)
	}
	if !strings.Contains(response.Body.String(), `"code":"NOT_FOUND"`) {
		t.Fatalf("expected JSON error code, got %s", response.Body.String())
	}
}
