package loaner

import "time"

type PaymentType string

const (
	PaymentTypeLoan      = PaymentType("loan")
	PaymentTypeRepayment = PaymentType("repayment")
)

type TransactionType string

const (
	CreditTransactionType = TransactionType("credit")
	DebitTransactionType  = TransactionType("debit")
)

type Transaction struct {
	ID              string          `json:"id,omitempty"`
	PaymentType     PaymentType     `json:"payment_type"`
	TransactionType TransactionType `json:"transaction_type"`
	Amount          float64         `json:"amount"`
	CreatedAt       time.Time       `json:"created_at,omitempty"`
	UpdatedAt       time.Time       `json:"updated_at,omitempty"`
}

type LoanService interface {
	InitiateLoan(trx *Transaction) error
	Repay(trx *Transaction) error
	SetRate(rate float64) error
	GetRate() float64
	Clear() error
	Transactions() []*Transaction
}
