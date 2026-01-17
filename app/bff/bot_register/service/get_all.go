package service

import (
	"bnbot/app/bff/bot_register/domain"
	"context"
	"errors"
)

func (s *service) GetAll(ctx context.Context) ([]*domain.BotRegister, error) {
	templates, err := s.botRegisterTemplateService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if templates == nil {
		return nil, errors.New("bot templates not found")
	}
	return templates, nil
}
