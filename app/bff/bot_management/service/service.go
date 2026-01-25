package service

import (
	"bnbot/app/bff/bot_management/domain"
	externalservice "bnbot/app/bff/bot_management/infrastructure/external_service"
	"context"
)

type IService interface {
	Get(ctx context.Context, botId string) (*domain.Bot, error)
	GetAll(ctx context.Context) ([]*domain.Bot, error)
	Insert(ctx context.Context, b *domain.Bot) error
	Update(ctx context.Context, b *domain.Bot) error
	Delete(ctx context.Context, botId string) error
}

type service struct {
	botService externalservice.IBotManagementService
}

func NewService(botService externalservice.IBotManagementService) IService {
	return &service{botService: botService}
}
