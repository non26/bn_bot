package db

import (
	"bnbot/app/core/bot_register_template/domain"
	"context"
)

type IBotBNTemplateRepository interface {
	Get(ctx context.Context, exchangeId string, templateId string) (*domain.BotTemplate, error)
	GetAll(ctx context.Context) ([]*domain.BotTemplate, error)
	Upsert(ctx context.Context, b *domain.BotTemplate) error
	Delete(ctx context.Context, exchangeId string, templateId string) error
}
