# Loaner API

There are two ways to run the application, the dockerised version or the local version

## Docker

`docker build -t loaner .`

`docker run -p 8888:8888 loaner`


- [Installation](#installation)
- [Starting App](#starting-app)

## Local Installation

- Make sure that [Golang](https://golang.org/) is installed, the minimum required is currently 1.12.5.

## Starting App

- Start the app with `make local-app`:

```shell
make local-app
```

You should see something similiar:

```shell
2019/12/02 12:30:02 connect to http://localhost:8888/ for GraphQL playground
```

- Navigate to [localhost:8080](http://localhost:8888). Please note `8080` is the default port used, you can change this in `main.go` or Dockerfile for your docker deployment

This should load up a graphql client you can use to test out the endpoints. Click on docs to see the graphql documentation.

Example Request and Response 

## InitiateLoan
```
mutation initiateLoanMutation($input: NewLoan!) {
  initiateLoan(input: $input)
}
```

**Query Variables**
```
{
  "input": {
    "amount": 5000,
    "rate": 0.1,
    "start": "2019-01-02T15:04:05Z"
  }
}
```

## Add Payment
```
mutation addPaymentMutation($amount: Float!, $date: Time!) {
  addPayment(amount: $amount, date: $date)
}
```

**Query Variables**
```
{
    "amount": 3000.06301369863,
    "date": "2019-01-06T15:04:05Z"
}
```

## Balance
```
query balance($date: Time) {
  balance(date: $date)
}
```

**Query Variables**
```
{
  "date": "2019-01-07T15:04:05Z"
}
```