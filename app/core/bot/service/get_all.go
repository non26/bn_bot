package service

import (
	"bnbot/app/core/bot/domain"
	"context"
)

func (s *service) GetAll(ctx context.Context) ([]*domain.Bot, error) {
	return s.repository.GetAll(ctx)
}
