package transaction

const maxRetryAttempts = 3

func CanRetryTransaction(status string, attempts int) bool {
	if status == "FAILED" {
		return false
	}
	if attempts < 0 {
		return false
	}
	if attempts >= maxRetryAttempts {
		return false
	}
	return true
}
