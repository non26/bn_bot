package db

import (
	"bnbot/app/core/bot_opening/domain"
	"context"
)

type IBotOnRunRepository interface {
	Get(ctx context.Context, botId string) (*domain.BotOpening, error)
	GetAll(ctx context.Context) ([]*domain.BotOpening, error)
	Upsert(ctx context.Context, b *domain.BotOpening) error
	Delete(ctx context.Context, botId string) error
}
