package validation

import (
	"errors"
	"strings"
)

func GenerateTransactionReference(userID string) (string, error) {
	if strings.TrimSpace(userID) == "" {
		return "", errors.New("userID cannot be empty")
	}

	ref := "TXN-" + userID
	return ref, nil
}
