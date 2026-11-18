package validation

import (
	"errors"
)

func ValidateWithdrawal(amount int64) error {
	if amount == 0 {
		return errors.New("invalid amount")
	}
	if amount < 0 {
		return errors.New("Amount cannot be negative!")
	}
	if amount < 500 {
		return errors.New("amount below min threshhold!")
	}
	if amount > 500000 {
		return errors.New("amount exceed maxim threshhold!")
	}
	return nil
}

func CheckSufficientBalance(balance, amount int64) error {
	if balance < 0 {
		return errors.New("balance cannot be negative")
	}
	if amount < 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if amount > balance {
		return errors.New("withdraw amount cant exceed available balance")
	}
	return nil
}
