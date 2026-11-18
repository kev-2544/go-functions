package payment

import (
	"errors"
	"fmt"
)

func ProcessPayment(amount, balance int64) (string, error) {
	if amount < 0 {
		return "", errors.New("Amount must be positive")
	}

	if amount < 100 {
		return "", errors.New("Minimum amount allowed is ksh 100")
	}

	fee := amount * 2 / 100
	total := amount + fee

	if total > balance {
		return "", errors.New("insufficient balance")
	}

	newBalance := balance - (amount + fee)

	return fmt.Sprintf("payment successful, new balance: %d", newBalance), nil
}
