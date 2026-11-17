package main

import(
	"fmt"
	"errors"
)

func validateWithdrawal(amount int64) error {
	if amount == 0 {
		return errors.New("invalid amount")
	}
	if amount <0 {
		return errors.New("Amount cannot be negative!")
	}
	if amount <  500 {
		return errors.New("amount below min threshhold!")
	}
	if amount > 500000 {
		return  errors.New("amount exceed maxim threshhold!")
	}
	return nil

}

func main (){
	amount:= int64(1000000)

	err := validateWithdrawal(amount)

	if err != nil {
		fmt.Println("Withdrawal declined:", err)
		return
	}

}