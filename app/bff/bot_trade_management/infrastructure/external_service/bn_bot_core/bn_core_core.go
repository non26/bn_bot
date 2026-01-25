package externalservice

import (
	"bnbot/app/bff/bot_trade_management/domain"
	"context"
)

type IExternalBnBotOpeningService interface {
	Get(ctx context.Context, botId string) (*domain.Trade, error)
	GetAll(ctx context.Context) ([]*domain.Trade, error)
	Insert(ctx context.Context, b *domain.Trade) error
	Update(ctx context.Context, b *domain.Trade) error
	Delete(ctx context.Context, botId string) error
}

type IExternalBotService interface {
	Get(ctx context.Context, botId string) (*domain.Trade, error)
	GetAll(ctx context.Context) ([]*domain.Trade, error)
	// Insert(ctx context.Context, b *domain.Trade) error
	// Update(ctx context.Context, b *domain.Trade) error
	// Delete(ctx context.Context, botId string) error
}
