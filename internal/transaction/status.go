package transaction

import (
	"errors"
	"strings"
)

func ValidateTransactionStatus(status string) (string, error) {
	status = strings.TrimSpace(status)
	status = strings.ToUpper(status)

	if status == "PENDING" || status == "PROCESSING" || status == "SUCCESS" || status == "FAILED" {
		return "status", nil
	}

	return "", errors.New("unknown status")
}
