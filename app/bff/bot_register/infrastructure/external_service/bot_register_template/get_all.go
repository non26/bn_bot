package externalservice

import (
	"bnbot/app/bff/bot_register/domain"
	"bnbot/app/bff/bot_register/infrastructure/external_service/bot_register_template/dto"
	"context"
)

func (s *botRegisterTemplateService) GetAll(ctx context.Context) ([]*domain.BotRegister, error) {
	templates, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if templates == nil {
		return nil, nil
	}
	res := []*domain.BotRegister{}
	for _, template := range templates {
		_res := &dto.BotRegisterTemplateDTO{}
		_res = _res.FromExternalServiceDomain(template)
		res = append(res, _res.ToDomain())
	}
	return res, nil
}
