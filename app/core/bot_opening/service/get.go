package service

import (
	"bnbot/app/core/bot_opening/domain"
	"context"
)

func (s *service) Get(ctx context.Context, botId string) (*domain.BotOpening, error) {
	return s.repository.Get(ctx, botId)
}
