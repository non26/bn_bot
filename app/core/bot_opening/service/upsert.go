package service

import (
	"bnbot/app/core/bot_opening/domain"
	"context"
)

func (s *service) Upsert(ctx context.Context, b *domain.BotOpening) error {
	return s.repository.Upsert(ctx, b)
}
