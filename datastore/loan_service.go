package datastore

import (
	"loaner"
)

type LoanService struct {
	transactions []*loaner.Transaction
	rate         float64
}

func NewLoanService(transactions []*loaner.Transaction) *LoanService {
	return &LoanService{transactions: transactions}
}

func (ls *LoanService) InitiateLoan(trx *loaner.Transaction) error {
	ls.transactions = append(ls.transactions, trx)
	return nil
}

func (ls *LoanService) Repay(trx *loaner.Transaction) error {
	ls.transactions = append(ls.transactions, trx)
	return nil
}

func (ls *LoanService) SetRate(rate float64) error {
	ls.rate = rate
	return nil
}

func (ls *LoanService) GetRate() float64 {
	return ls.rate
}

func (ls *LoanService) Clear() error {
	ls.transactions = []*loaner.Transaction{}
	return nil
}

func (ls *LoanService) Transactions() []*loaner.Transaction {
	return ls.transactions
}
