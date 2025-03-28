package postgres

import (
	"github.com/shopspring/decimal"
	"gw-exchanger/internal/storages"
	"gw-exchanger/pkg/client/postgresql"
	"gw-exchanger/pkg/logger"
)

type currencyRepository struct{}

func NewCurrencyRepository() storages.CurrencyRepository {
	return &currencyRepository{}
}

func (r *currencyRepository) GetExchangeRates() ([]storages.Exchange, error) {
	logger.Log.Debug("Trying get exchange rates")
	db, err := postgresql.NewClient()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM exchange_rates")
	if err != nil {
		return nil, err
	}

	var exchanges []storages.Exchange

	for rows.Next() {
		var exchange storages.Exchange
		if err := rows.Scan(&exchange.Currency, &exchange.Rate); err != nil {
			return exchanges, err
		}
		exchanges = append(exchanges, exchange)
	}
	if err = rows.Err(); err != nil {
		return exchanges, err
	}

	logger.Log.Info("Successfully got exchanges", "Exchanges", exchanges)

	return exchanges, nil
}

func (r *currencyRepository) GetExchangeRateForCurrency(from, to storages.Currency) (decimal.Decimal, error) {
	logger.Log.Debug("Trying get exchange rate currency", "from", from, "to", to)
	db, err := postgresql.NewClient()
	if err != nil {
		return decimal.Decimal{}, err
	}
	defer db.Close()

	var rate decimal.Decimal

	result := db.QueryRow("SELECT ((SELECT rate FROM exchange_rates WHERE currency = $1) / (SELECT rate FROM exchange_rates WHERE currency = $2));", from, to)
	err = result.Scan(&rate)
	if err != nil {
		return decimal.Decimal{}, err
	}

	logger.Log.Info("Successfully got exchange", "from", from, "to", to, "rate", rate)

	return rate, nil
}
