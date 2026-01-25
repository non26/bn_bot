package externalservice

import (
	"bnbot/app/bff/bot_trade_management/domain"
	"context"
)

type IExternalBnBffTradeService interface {
	NewOrder(ctx context.Context, d_req *domain.Trade) error
}

type IExternalBnBffPositionManagementService interface {
	GetCurrentPosition(ctx context.Context, d_req *domain.Trade) (*domain.Trade, error)
}
