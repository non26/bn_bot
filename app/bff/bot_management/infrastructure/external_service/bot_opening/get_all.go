package externalservice

import (
	"bnbot/app/bff/bot_management/domain"
	"bnbot/app/bff/bot_management/infrastructure/external_service/bot_opening/dto"
	"context"
)

func (s *externalBotOpeningService) GetAll(ctx context.Context) ([]*domain.BotOpening, error) {
	botOpenings, err := s.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if botOpenings == nil {
		return nil, nil
	}
	res := []*domain.BotOpening{}
	for _, botOpening := range botOpenings {
		botOpeningDTO := dto.NewEmptyBotOpening()
		res = append(res, botOpeningDTO.FromExternalBotOpeningServiceDomainToDomain(botOpening))
	}
	return res, nil
}
