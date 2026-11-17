package main

import (
	"fmt"

)

const maxRetryAttempts = 3

func canRetryTransaction(status string, attempts int) bool {
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

func main(){
	status := "FAILED"
	attempts := 0
	allowed := canRetryTransaction(status, attempts)

	fmt.Println("retry allowed", allowed)
}