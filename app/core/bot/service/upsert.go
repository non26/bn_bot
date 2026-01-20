package service

import (
	"bnbot/app/core/bot/domain"
	"context"
)

func (s *service) Upsert(ctx context.Context, b *domain.Bot) error {
	return s.repository.Upsert(ctx, b)
}
