package payment

import (
	"errors"
)

func CalculateTransactionFee(amount int64) (int64, error) {
	if amount <= 0 {
		return 0, errors.New("invalid amount")
	}
	fee := amount * 2 / 100

	if fee > 1000 {
		fee = 1000
	}
	return fee, nil
}
