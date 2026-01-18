package service

import (
	"bnbot/app/bff/market_data/domain"
	"context"
)

func (s *service) GetKline(ctx context.Context, d *domain.Kline) ([]*domain.Kline, error) {
	return s.marketDataService.GetKline(ctx, d)
}
