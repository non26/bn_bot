package service

import (
	"bnbot/app/bff/bot_management/domain"
	"context"
)

func (s *service) GetAll(ctx context.Context) ([]*domain.Bot, error) {
	bots, err := s.botService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return bots, nil
}
