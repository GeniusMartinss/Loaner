//go:generate go run github.com/99designs/gqlgen -v

package graphql

import (
	"context"
	"loaner/features/loan"
	"time"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	LoanHandler *loan.Handler
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) InitiateLoan(ctx context.Context, input NewLoan) (bool, error) {
	err := r.LoanHandler.InitiateLoan(input.Amount, input.Rate, input.Start)
	return err == nil, err
}

func (r *mutationResolver) AddPayment(ctx context.Context, amount float64, date time.Time) (bool, error) {
	err := r.LoanHandler.AddPayment(amount, date)
	return err == nil, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Balance(ctx context.Context, date *time.Time) (float64, error) {
	return r.LoanHandler.Balance(*date)
}
