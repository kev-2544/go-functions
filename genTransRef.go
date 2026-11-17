package main

import (
	"fmt"
	"errors"
	"strings"
)

func generateTransactionReference(userID string) (string, error) {
	if strings.TrimSpace(userID) == "" {
		return "", errors.New("userID cannot be empty")

	}
	
	ref := "TXN-"+ userID
	return ref, nil
}

func main(){
	userID := "PAYD-USERID-12345"

	ref, err := generateTransactionReference(userID)

	if err != nil {
		fmt.Println("failed to gen transaction reference:", err)
		return
	}
	fmt.Println("Reference:", ref)
}