package storages

import "github.com/shopspring/decimal"

type CurrencyRepository interface {
	GetExchangeRates() ([]Exchange, error)
	GetExchangeRateForCurrency(from, to Currency) (decimal.Decimal, error)
}
