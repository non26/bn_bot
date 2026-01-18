package service

import (
	"bnbot/app/core/bot_opening/domain"
	"context"
)

func (s *service) GetAll(ctx context.Context) ([]*domain.BotOpening, error) {
	return s.repository.GetAll(ctx)
}
