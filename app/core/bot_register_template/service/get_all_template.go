package service

import (
	"bnbot/app/core/bot_register_template/domain"
	"context"
)

func (s *botBNTemplateService) GetAll(ctx context.Context) ([]*domain.BotTemplate, error) {
	return s.repository.GetAll(ctx)
}
