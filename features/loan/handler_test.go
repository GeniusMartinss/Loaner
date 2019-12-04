package loan

import (
	"errors"
	"loaner"
	"loaner/datastore"
	"testing"
	"time"
)

func TestHandler_InitiateLoan(t *testing.T) {
	loanService := datastore.NewLoanService([]*loaner.Transaction{})
	loanHandler := NewHandler(loanService)

	loanAmount := 560.34
	rate := 0.34
	start := time.Now()

	if err := loanHandler.InitiateLoan(loanAmount, rate, start); err != nil {
		t.Errorf("InitiateLoan:, expected %v got %v", nil, err)
	}

	if setRate := loanService.GetRate(); setRate != rate {
		t.Errorf("InitiateLoan Rate:, expected %v got %v", rate, setRate)
	}
}

func TestHandler_AddPayment(t *testing.T) {
	loanService := datastore.NewLoanService([]*loaner.Transaction{})
	loanHandler := NewHandler(loanService)

	loanAmount := 560.34
	rate := 0.34
	start := time.Now()

	if err := loanHandler.InitiateLoan(loanAmount, rate, start); err != nil {
		t.Errorf("InitiateLoan:, expected %v got %v", nil, err)
	}

	if err := loanHandler.AddPayment(loanAmount, start); err != nil {
		t.Errorf("AddPayment:, expected %v got %v", nil, err)
	}
}

func TestHandler_AddPayment_ShouldFail_If_No_Loan(t *testing.T) {
	loanService := datastore.NewLoanService([]*loaner.Transaction{})
	loanHandler := NewHandler(loanService)

	loanAmount := 560.34
	start := time.Now()

	if err := loanHandler.AddPayment(loanAmount, start); err == nil {
		t.Errorf("AddPayment:, expected %v got %v", errors.New("cannot make payment without an unpaid loan"), err)
	}
}

func TestHandler_Balance_ShouldBe_LoanAmount_At_Date_OfLoan(t *testing.T) {
	loanService := datastore.NewLoanService([]*loaner.Transaction{})
	loanHandler := NewHandler(loanService)

	loanAmount := 560.34
	rate := 0.34
	start := time.Now()

	if err := loanHandler.InitiateLoan(loanAmount, rate, start); err != nil {
		t.Errorf("InitiateLoan:, expected %v got %v", nil, err)
	}

	balance, err := loanHandler.Balance(start)
	if err != nil {
		t.Errorf("Balance:, expected %v got %v", nil, err)
	}

	if balance != loanAmount {
		t.Errorf("Balance:, expected %v got %v", loanAmount, balance)
	}
}

// Payment is made on the same day as the loan, so no interest yet
func TestHandler_Balance_Should_Be_Zero_If_Payment_Of_LoanAmount_Is_Made(t *testing.T) {
	loanService := datastore.NewLoanService([]*loaner.Transaction{})
	loanHandler := NewHandler(loanService)

	loanAmount := 560.34
	rate := 0.34
	start := time.Now()

	if err := loanHandler.InitiateLoan(loanAmount, rate, start); err != nil {
		t.Errorf("InitiateLoan:, expected %v got %v", nil, err)
	}

	if err := loanHandler.AddPayment(loanAmount, start); err != nil {
		t.Errorf("AddPayment:, expected %v got %v", nil, err)
	}

	balance, err := loanHandler.Balance(start)
	if err != nil {
		t.Errorf("Balance:, expected %v got %v", nil, err)
	}

	if balance != float64(0) {
		t.Errorf("Balance:, expected %v got %v", 0, balance)
	}
}
