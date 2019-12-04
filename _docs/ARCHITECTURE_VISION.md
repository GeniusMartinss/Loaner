# Architecture Vision

The code structure holds a top level interface ```LoanService``` that is the building block for the entire application
This allows any datastore that implements this interface to be suitable as a data layer.

## Features
Each feature lives in its own package and has everything it needs to work there (logic, data access).
For now the only feature is Loan, if Interest become a transaction on their own, they will live as a separate feature.

## Handlers
Each feature has a handler for interacting with the storage layer. This is where the majority of the business logic sits
Handler has a unit test for all the functionality it handles.

## Datastore
The core of the application is a ```Transaction``` which is the data structure that represents a loan or a loan payment
A Loan object is basically a structure holding the rate set at loan initiation and a list of transactions, with the genesis transaction being the initial loan.


```
type LoanService struct {
   	transactions []*loaner.Transaction
   	rate		float64
   }
```
The design was made this way to keep a log of every payment with a dependency on the original loan which is a transaction, but a debit transaction.

Two types are used to distinguish a loan transaction from a repayment transaction

```	
    CreditTransactionType = TransactionType("credit")
   	DebitTransactionType  = TransactionType("debit")
```
Credit is a request that increases the value of the balance. e.g InitiateLoan

Debit is a request that reduces the value of the balance, e.g Loan repayment

```
	PaymentTypeLoan      = PaymentType("loan")
	PaymentTypeRepayment = PaymentType("repayment")
```

