package externalservice

import (
	"bnbot/app/bff/bot_register/domain"
	"context"
)

type IBotRegisterTemplateService interface {
	Get(ctx context.Context, exchangeId string, templateId string) (*domain.BotRegister, error)
	GetAll(ctx context.Context) ([]*domain.BotRegister, error)
	Upsert(ctx context.Context, b *domain.BotRegister) error
	Delete(ctx context.Context, exchangeId string, templateId string) error
}
