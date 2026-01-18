package externalservice

import (
	"bnbot/app/bff/market_data/domain"
	"context"
)

type IMarketDataService interface {
	GetKline(ctx context.Context, req *domain.Kline) ([]*domain.Kline, error)
	GetPreviousKline(ctx context.Context, req *domain.Kline) (*domain.Kline, error)
}
