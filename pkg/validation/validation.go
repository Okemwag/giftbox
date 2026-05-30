package validation

import (
	"errors"
	"net/mail"
	"strings"
)

func Required(name, value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New(name + " is required")
	}
	return nil
}

func Email(value string) error {
	if _, err := mail.ParseAddress(value); err != nil {
		return errors.New("invalid email address")
	}
	return nil
}

func KenyanPhone(value string) error {
	normalized := strings.TrimSpace(value)
	if strings.HasPrefix(normalized, "+254") && len(normalized) == 13 {
		return nil
	}
	if strings.HasPrefix(normalized, "254") && len(normalized) == 12 {
		return nil
	}
	return errors.New("phone must be in +254 or 254 format")
}
