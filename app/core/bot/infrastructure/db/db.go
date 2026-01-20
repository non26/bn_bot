package db

import (
	"bnbot/app/core/bot/domain"
	"context"
)

type IBotRepository interface {
	Get(ctx context.Context, botId string) (*domain.Bot, error)
	GetAll(ctx context.Context) ([]*domain.Bot, error)
	Upsert(ctx context.Context, b *domain.Bot) error
	Delete(ctx context.Context, botId string) error
}
