package main

import (
	"fmt"
	"errors"
	"strings"
)

func validateTransactionStatus(status string) (string, error) {

	status = strings.TrimSpace(status)
	status = strings.ToUpper(status)



	if status == "PENDING" || status == "PROCESSING" || status == "SUCCESS" || status == "FAILED" {
		return "status", nil
	}

	return "", errors.New("unknown status")
	

}

func main () {
	status:= "  jjpending "
	
	status, err:= validateTransactionStatus(status)

	if err != nil {
		fmt.Println("status declined:", err)
		return
	}
}
