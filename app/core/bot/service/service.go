package service

import (
	"bnbot/app/core/bot/domain"
	"bnbot/app/core/bot/infrastructure/db"
	"context"
)

type IService interface {
	Get(ctx context.Context, botId string) (*domain.Bot, error)
	GetAll(ctx context.Context) ([]*domain.Bot, error)
	Upsert(ctx context.Context, b *domain.Bot) error
	Delete(ctx context.Context, botId string) error
}

type service struct {
	repository db.IBotRepository
}

func NewService(repository db.IBotRepository) IService {
	return &service{repository: repository}
}
