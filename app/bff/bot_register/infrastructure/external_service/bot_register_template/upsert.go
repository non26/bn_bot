package externalservice

import (
	"bnbot/app/bff/bot_register/domain"
	"bnbot/app/bff/bot_register/infrastructure/external_service/bot_register_template/dto"
	"context"
)

func (s *botRegisterTemplateService) Upsert(ctx context.Context, b *domain.BotRegister) error {
	req := &dto.BotRegisterTemplateDTO{}
	req = req.FromDomain(b)
	err := s.externalBotRegisterTemplatCoreService.Upsert(ctx, req.ToExternalServiceDomain())
	if err != nil {
		return err
	}

	return nil
}
