package service

import (
	"bnbot/app/core/bot_opening/domain"
	"bnbot/app/core/bot_opening/infrastructure/db"
	"context"
)

type IService interface {
	Get(ctx context.Context, botId string) (*domain.BotOpening, error)
	GetAll(ctx context.Context) ([]*domain.BotOpening, error)
	Upsert(ctx context.Context, b *domain.BotOpening) error
	Delete(ctx context.Context, botId string) error
}

type service struct {
	repository db.IBotOnRunRepository
}

func NewService(repository db.IBotOnRunRepository) IService {
	return &service{repository: repository}
}
