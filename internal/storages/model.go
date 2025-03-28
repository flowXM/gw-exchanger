package storages

import "github.com/shopspring/decimal"

type Exchange struct {
	Currency Currency
	Rate     decimal.Decimal
}

type Currency string

const (
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)
