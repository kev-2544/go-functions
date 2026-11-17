package main

import (
	"fmt"
	"errors"
	"strings"
)

func processWebhook(reference, status string) error {
	reference =strings.TrimSpace(reference)
	status =strings.TrimSpace(status)

	if !strings.HasPrefix(reference, "PAYD-") {
		return errors.New("Reference must start with PAYD-")
	}
	if status != "PENDING" && status != "SUCCESS" && status != "FAILED" {
		return errors.New("unsupported status")

	}
	return nil

}

func main(){
	reference := "PAYD-1345"
	status := "FAILED"

	err := processWebhook(reference, status)

	if err != nil {
		fmt.Println("webhook processing failed:", err)
		return
	}
	

}