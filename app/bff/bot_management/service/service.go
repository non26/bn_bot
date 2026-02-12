package service

import (
	"bnbot/app/bff/bot_management/domain"
	"context"
)

type IBotService interface {
	Get(ctx context.Context, botId string) (*domain.Bot, error)
	GetAll(ctx context.Context) ([]*domain.Bot, error)
	Insert(ctx context.Context, b *domain.Bot) error
	Update(ctx context.Context, b *domain.Bot) error
	Delete(ctx context.Context, botId string) error
}

type IBotOpeningService interface {
	Get(ctx context.Context, botId string) (*domain.BotOpening, error)
	GetAll(ctx context.Context) ([]*domain.BotOpening, error)
	Insert(ctx context.Context, b *domain.BotOpening) error
	Update(ctx context.Context, b *domain.BotOpening) error
	Delete(ctx context.Context, botId string) error
}
