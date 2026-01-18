package service

import (
	"bnbot/app/bff/market_data/domain"
	"context"
)

func (s *service) GetPreviousKline(ctx context.Context, d *domain.Kline) (*domain.Kline, error) {
	return s.marketDataService.GetPreviousKline(ctx, d)
}
