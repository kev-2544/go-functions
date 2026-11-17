package main

import (
	"fmt"
	"errors"
)

func processPayment(amount, balance int64) (string, error) {
	if amount < 0 {
		return "", errors.New("Amount must be positive")
	}
	
	if amount < 100 {
		return "", errors.New("Minimum amount allowed is ksh 100")
	}

	fee := amount * 2/100
	total := amount + fee

	if total > balance {
		return "", errors.New("insufficient balance")
	}
	
	
	 
	
	newBalance := balance - (amount + fee)

	return fmt.Sprintf("payment successful, new balance: %d", newBalance), nil

	
}

func main(){
	amount := int64(1000)
	balance := int64(1500)

	result, err := processPayment(amount, balance)

	if err != nil {
		fmt.Println("payment failed:", err)
		return
	}
	fmt.Println("status:", result)

}