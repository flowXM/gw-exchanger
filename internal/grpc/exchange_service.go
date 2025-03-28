package grpc

import (
	"context"
	pe "github.com/flowXM/proto-exchange/exchange"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gw-exchanger/internal/storages"
	"gw-exchanger/internal/storages/postgres"
	"gw-exchanger/pkg/logger"
)

type ExchangeServiceServer struct {
	pe.UnimplementedExchangeServiceServer
}

func (e *ExchangeServiceServer) GetExchangeRates(context.Context, *pe.Empty) (*pe.ExchangeRatesResponse, error) {
	cr := postgres.NewCurrencyRepository()
	exchangeRates, err := cr.GetExchangeRates()
	if err != nil {
		logger.Log.Error("gRPC", "error", err)
		return &pe.ExchangeRatesResponse{}, status.Error(codes.Internal, "")
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
		logger.Log.Error("gRPC", "error", err)
		return &pe.ExchangeRateResponse{}, status.Error(codes.Internal, "")
	}

	return &pe.ExchangeRateResponse{
		FromCurrency: request.FromCurrency,
		ToCurrency:   request.ToCurrency,
		Rate:         float32(exchangeRate.InexactFloat64()),
	}, nil
}
