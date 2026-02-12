package service

import (
	"bnbot/app/bff/bot_management/domain"
	"context"
)

func (s *botOpeningService) GetAll(ctx context.Context) ([]*domain.BotOpening, error) {
	return s.externalService.GetAll(ctx)
}
