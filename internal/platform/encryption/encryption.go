package encryption

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func HashSecret(value string) string {
	sum := sha256.Sum256([]byte(value))
	return hex.EncodeToString(sum[:])
}

func Redact(value string) string {
	value = strings.TrimSpace(value)
	if len(value) <= 4 {
		return "****"
	}
	return strings.Repeat("*", len(value)-4) + value[len(value)-4:]
}
