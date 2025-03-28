package grpc

import (
	"context"
	pe "github.com/flowXM/proto-exchange/exchange"
	"gw-exchanger/internal/storages"
	"gw-exchanger/internal/storages/postgres"
)

type ExchangeServiceServer struct {
	pe.UnimplementedExchangeServiceServer
}

func (e *ExchangeServiceServer) GetExchangeRates(context.Context, *pe.Empty) (*pe.ExchangeRatesResponse, error) {
	cr := postgres.NewCurrencyRepository()
	exchangeRates, err := cr.GetExchangeRates()
	if err != nil {
		return nil, err
	}
	var mp = make(map[string]float32)

	for _, rate := range exchangeRates {
		mp[string(rate.Currency)] = float32(rate.Rate.InexactFloat64())
	}

	return &pe.ExchangeRatesResponse{
		Rates: mp,
	}, nil
}

func (e *ExchangeServiceServer) GetExchangeRateForCurrency(ctx context.Context, request *pe.CurrencyRequest) (*pe.ExchangeRateResponse, error) {
	cr := postgres.NewCurrencyRepository()
	exchangeRate, err := cr.GetExchangeRateForCurrency(storages.Currency(request.FromCurrency), storages.Currency(request.ToCurrency))
	if err != nil {
		return nil, err
	}

	return &pe.ExchangeRateResponse{
		FromCurrency: request.FromCurrency,
		ToCurrency:   request.ToCurrency,
		Rate:         float32(exchangeRate.InexactFloat64()),
	}, nil
}
