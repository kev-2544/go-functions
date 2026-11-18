package main

import (
	"fmt"

	"functions/internal/payment"
	"functions/internal/transaction"
	"functions/internal/validation"
)

func main() {
	// ----- validation -----

	err := validation.ValidateWithdrawal(1000)
	fmt.Println("ValidateWithdrawal(1000):", err)

	err = validation.CheckSufficientBalance(2000, 1900)
	fmt.Println("CheckSufficientBalance(2000, 1900):", err)

	err = validation.ValidatePhoneNumber("254322222222")
	fmt.Println("ValidatePhoneNumber(254322222222):", err)

	ref, err := validation.GenerateTransactionReference("PAYD-USERID-12345")
	fmt.Println("GenerateTransactionReference:", ref, err)

	// ----- payment -----

	fee, err := payment.CalculateTransactionFee(90000)
	fmt.Println("CalculateTransactionFee(90000):", fee, err)

	result, err := payment.ProcessPayment(1000, 1500)
	fmt.Println("ProcessPayment(1000, 1500):", result, err)

	allowed := payment.CanTransitionPayment("PENDING", "PROCESSING")
	fmt.Println("CanTransitionPayment(PENDING -> PROCESSING):", allowed)

	// ----- transaction -----

	allowed = transaction.CanRetryTransaction("FAILED", 0)
	fmt.Println("CanRetryTransaction(FAILED, 0):", allowed)

	err = transaction.ProcessWebhook("PAYD-1345", "FAILED")
	fmt.Println("ProcessWebhook(PAYD-1345, FAILED):", err)

	status, err := transaction.ValidateTransactionStatus("  pending  ")
	fmt.Println("ValidateTransactionStatus(\"  pending  \"):", status, err)
}
