package main


import (
	"fmt"
	"errors"
)

func checkSufficientBalance(balance, amount int64) error {
	if balance < 0{
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

func main(){
	balance := int64(2000)
	amount := int64(1900)

	err := checkSufficientBalance(balance, amount)

	if err != nil {
		fmt.Println("withdrawaal failed:", err)
		return
	}
}