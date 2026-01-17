package service

import (
	"bnbot/app/bff/bot_register/domain"
	externalservice "bnbot/app/bff/bot_register/infrastructure/external_service"
	"context"
)

type IService interface {
	Get(ctx context.Context, exchangeId string, templateId string) (*domain.BotRegister, error)
	GetAll(ctx context.Context) ([]*domain.BotRegister, error)
	Delete(ctx context.Context, exchangeId string, templateId string) error
	Set(ctx context.Context, b *domain.BotRegister) error
	Update(ctx context.Context, b *domain.BotRegister) error
}

type service struct {
	botRegisterTemplateService externalservice.IBotRegisterTemplateService
}

func NewService(botRegisterTemplateService externalservice.IBotRegisterTemplateService) IService {
	return &service{botRegisterTemplateService: botRegisterTemplateService}
}
