package transaction

import (
	"errors"
	"strings"
)

func ProcessWebhook(reference, status string) error {
	reference = strings.TrimSpace(reference)
	status = strings.TrimSpace(status)

	if !strings.HasPrefix(reference, "PAYD-") {
		return errors.New("Reference must start with PAYD-")
	}
	if status != "PENDING" && status != "SUCCESS" && status != "FAILED" {
		return errors.New("unsupported status")
	}
	return nil
}
