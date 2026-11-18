# go-functions

Practice Go functions for payment and transaction checks.

## Folder layout

| Folder | What's in it |
|---|---|
| `cmd/main.go` | The program you run. It calls every function with test values. |
| `internal/validation` | Withdrawal amount, balance, phone number, and transaction reference checks |
| `internal/payment` | Fee calculation, processing a payment, and status transitions |
| `internal/transaction` | Retry rules, webhook checks, and status validation |

## Run it

From this folder:

```
go run ./cmd
```
