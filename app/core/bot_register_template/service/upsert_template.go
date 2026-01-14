package service

import (
	"bnbot/app/core/bot_register_template/domain"
	"context"
)

func (s *botBNTemplateService) Upsert(ctx context.Context, b *domain.BotTemplate) error {
	return s.repository.Upsert(ctx, b)
}
