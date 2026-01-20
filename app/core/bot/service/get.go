package service

import (
	"bnbot/app/core/bot/domain"
	"context"
)

func (s *service) Get(ctx context.Context, botId string) (*domain.Bot, error) {
	return s.repository.Get(ctx, botId)
}
