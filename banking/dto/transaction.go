package dto

import "github.com/moisotico/banking/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	// return true or false
	return r.TransactionType == WITHDRAWAL
}
func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	// return true or false
	return r.TransactionType == DEPOSIT
}

func (r TransactionRequest) Validate() *errs.AppError {
	if !r.IsTransactionTypeWithdrawal() && !r.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction can only be a deposit or withdrawal")
	}
	if r.Amount <= 0 {
		return errs.NewValidationError("Amount must be positive, and more than zero")
	}
	return nil
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
