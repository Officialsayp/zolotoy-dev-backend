package domain

import "errors"

type Money struct {
	Amount   int64
	Currency Currency
}

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyRUB Currency = "RUB"
	CurrencyRSD Currency = "RSD"
)

func NewMoney(amount int64, currency Currency) (Money, error) {
	switch {
	case amount < 0:
		return Money{}, errors.New("amount must be greater than zero")
	case currency != CurrencyUSD && currency != CurrencyEUR && currency != CurrencyRUB && currency != CurrencyRSD:
		return Money{}, errors.New("currency not supported")
	}
	return Money{
		Amount:   amount,
		Currency: currency,
	}, nil
}
