package externalservice

import (
	"bnbot/app/bff/bot_management/domain"
	"bnbot/app/bff/bot_management/infrastructure/external_service/bot/dto"
	"context"
)

func (s *botService) Insert(ctx context.Context, b *domain.Bot) error {
	reqdto := dto.NewEmptyBot()
	reqExternalService := reqdto.FromDomainToExternalBotServiceDomain(b)
	err := s.externalbotService.Upsert(ctx, reqExternalService)
	if err != nil {
		return err
	}
	return nil
}
