package service

import (
	"bnbot/app/bff/market_data/domain"
	externalservice "bnbot/app/bff/market_data/infrastructure/external_service"
	"context"
)

type IService interface {
	GetKline(ctx context.Context, d *domain.Kline) ([]*domain.Kline, error)
	GetPreviousKline(ctx context.Context, d *domain.Kline) (*domain.Kline, error)
}

type service struct {
	marketDataService externalservice.IMarketDataService
}

func NewService(marketDataService externalservice.IMarketDataService) IService {
	return &service{marketDataService: marketDataService}
}
