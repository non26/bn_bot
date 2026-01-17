package service

import (
	"bnbot/app/bff/bot_register/domain"
	"context"
	"errors"
)

func (s *service) Update(ctx context.Context, b *domain.BotRegister) error {
	item, err := s.botRegisterTemplateService.Get(ctx, b.ExchangeId, b.TemplateId)
	if err != nil {
		return err
	}
	if item == nil {
		return errors.New("bot template not found")
	}
	err = s.botRegisterTemplateService.Upsert(ctx, b)
	if err != nil {
		return err
	}
	return nil
}
