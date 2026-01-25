package externalservice

import (
	"bnbot/app/bff/bot_management/domain"
	"context"
)

type IBotManagementService interface {
	Get(ctx context.Context, botId string) (*domain.Bot, error)
	GetAll(ctx context.Context) ([]*domain.Bot, error)
	Insert(ctx context.Context, b *domain.Bot) error
	Update(ctx context.Context, b *domain.Bot) error
	Delete(ctx context.Context, botId string) error
}
