package main

import (
	"fmt"
	"errors"
)

func calculateTransactionFee(amount int64) (int64, error) {
	if amount <= 0 {
		return 0, errors.New("invalid amount")
	}
	fee := amount * 2/100

	if fee > 1000 {
		fee = 1000
	}
	return fee, nil

}

func main(){
	amount := int64(90000)

	fee, err := calculateTransactionFee(amount)

	if err != nil {
		fmt.Println("fee calculation failed:", err)
		return
	}
	fmt.Println("Transaction fee:", fee)

}