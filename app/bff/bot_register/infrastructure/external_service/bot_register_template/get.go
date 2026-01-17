package externalservice

import (
	"bnbot/app/bff/bot_register/domain"
	"bnbot/app/bff/bot_register/infrastructure/external_service/bot_register_template/dto"
	"context"
)

func (s *botRegisterTemplateService) Get(ctx context.Context, exchangeId string, templateId string) (*domain.BotRegister, error) {
	item, err := s.repository.Get(ctx, exchangeId, templateId)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, nil
	}
	res := &dto.BotRegisterTemplateDTO{}
	res = res.FromExternalServiceDomain(item)
	return res.ToDomain(), nil
}
