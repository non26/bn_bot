package service

import (
	"bnbot/app/bff/bot_register/domain"
	"context"
	"errors"
)

func (s *service) Get(ctx context.Context, exchangeId string, templateId string) (*domain.BotRegister, error) {
	item, err := s.botRegisterTemplateService.Get(ctx, exchangeId, templateId)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("bot template not found")
	}
	return item, nil
}
