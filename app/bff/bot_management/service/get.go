package service

import (
	"bnbot/app/bff/bot_management/domain"
	"context"
)

func (s *service) Get(ctx context.Context, botId string) (*domain.Bot, error) {
	bot, err := s.botService.Get(ctx, botId)
	if err != nil {
		return nil, err
	}
	if bot == nil {
		return nil, nil
	}
	return bot, nil
}
