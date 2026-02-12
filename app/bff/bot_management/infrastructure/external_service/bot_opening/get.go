package externalservice

import (
	"bnbot/app/bff/bot_management/domain"
	"bnbot/app/bff/bot_management/infrastructure/external_service/bot_opening/dto"
	"context"
)

func (s *externalBotOpeningService) Get(ctx context.Context, botId string) (*domain.BotOpening, error) {
	botOpening, err := s.service.Get(ctx, botId)
	if err != nil {
		return nil, err
	}
	if botOpening == nil {
		return nil, nil
	}
	botOpeningDTO := dto.NewEmptyBotOpening()
	res := botOpeningDTO.FromExternalBotOpeningServiceDomainToDomain(botOpening)
	return res, nil
}
