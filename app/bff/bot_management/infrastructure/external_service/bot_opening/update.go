package externalservice

import (
	"bnbot/app/bff/bot_management/domain"
	"bnbot/app/bff/bot_management/infrastructure/external_service/bot_opening/dto"
	"context"
)

func (s *externalBotOpeningService) Update(ctx context.Context, b *domain.BotOpening) error {
	reqdto := dto.NewEmptyBotOpening()
	reqExternalService := reqdto.FromDomainToExternalBotOpeningServiceDomain(b)
	err := s.service.Upsert(ctx, reqExternalService)
	if err != nil {
		return err
	}
	return nil
}
