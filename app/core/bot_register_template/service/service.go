package service

import (
	"bnbot/app/core/bot_register_template/domain"
	"bnbot/app/core/bot_register_template/infrastructure/db"
	"context"
)

type IBotBNTemplateService interface {
	Get(ctx context.Context, exchangeId string, templateId string) (*domain.BotTemplate, error)
	GetAll(ctx context.Context) ([]*domain.BotTemplate, error)
	Upsert(ctx context.Context, b *domain.BotTemplate) error
	Delete(ctx context.Context, exchangeId string, templateId string) error
}

type botBNTemplateService struct {
	repository db.IBotBNTemplateRepository
}

func NewBotBNTemplateService(repository db.IBotBNTemplateRepository) IBotBNTemplateService {
	return &botBNTemplateService{repository: repository}
}
