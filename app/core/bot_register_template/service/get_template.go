package service

import (
	"bnbot/app/core/bot_register_template/domain"
	"context"
)

func (s *botBNTemplateService) Get(ctx context.Context, exchangeId string, templateId string) (*domain.BotTemplate, error) {
	return s.repository.Get(ctx, exchangeId, templateId)
}
