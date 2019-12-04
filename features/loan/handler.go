package loan

import (
	"errors"
	"loaner"
	"log"
	"sort"
	"time"
)

type Handler struct {
	loanService loaner.LoanService
}

func NewHandler(loanService loaner.LoanService) *Handler {
	return &Handler{loanService: loanService}
}

func (lh *Handler) InitiateLoan(amount, rate float64, start time.Time) error {
	log.Printf("Initiating Loan of %f at a rate of %f", amount, rate)
	// Clear all previous loans.
	err := lh.loanService.Clear()
	if err != nil {
		return err
	}

	trx := &loaner.Transaction{
		ID:              "",
		PaymentType:     loaner.PaymentTypeLoan,
		TransactionType: loaner.CreditTransactionType,
		Amount:          amount,
		CreatedAt:       start,
		UpdatedAt:       time.Now(),
	}

	err = lh.loanService.SetRate(rate)
	if err != nil {
		return err
	}

	return lh.loanService.InitiateLoan(trx)
}

func (lh *Handler) AddPayment(amount float64, start time.Time) error {
	log.Printf("Making Payment of %f for date %v", amount, start)
	if len(lh.loanService.Transactions()) == 0 {
		return errors.New("cannot make payment without an unpaid loan")
	}
	trx := &loaner.Transaction{
		ID:              "",
		PaymentType:     loaner.PaymentTypeRepayment,
		TransactionType: loaner.DebitTransactionType,
		Amount:          amount,
		CreatedAt:       start,
		UpdatedAt:       time.Now(),
	}

	return lh.loanService.Repay(trx)
}

func (lh *Handler) Balance(date time.Time) (float64, error) {
	rate := lh.loanService.GetRate()
	transactions := lh.loanService.Transactions()

	if len(transactions) == 0 {
		return 0, nil
	}

	p := make(KeyValueList, len(transactions))
	i := 0
	for _, v := range transactions {
		p[i] = *v
		i++
	}

	sort.Sort(sort.Reverse(p))

	loanTrx := transactions[0]
	interest := 0.0
	principalBalance := loanTrx.Amount
	lastValidPayment := loanTrx.CreatedAt

	for i := 1; i < len(transactions); i++ {
		if transactions[i].CreatedAt.After(date) {
			break
		}
		// Time elapsed between payments
		interestPeriod := transactions[i].CreatedAt.Sub(transactions[i-1].CreatedAt).Hours() / 24
		interest += (rate / 100 / 365 * principalBalance) * interestPeriod
		principalBalance -= transactions[i].Amount
		lastValidPayment = transactions[i].CreatedAt
	}

	// If user still owes some money, and a future balance is requested
	projectedInterestPeriod := date.Sub(lastValidPayment).Hours() / 24
	interest += (rate / 100 / 365 * principalBalance) * projectedInterestPeriod

	log.Printf("Checking Balance for date %v", date)
	return principalBalance + interest, nil
}

type KeyValueList []loaner.Transaction

func (k KeyValueList) Less(i, j int) bool {
	return k[i].CreatedAt.Before(k[j].CreatedAt)
}

func (k KeyValueList) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k KeyValueList) Len() int {
	return len(k)
}
