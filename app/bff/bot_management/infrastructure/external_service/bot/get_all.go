package externalservice

import (
	"bnbot/app/bff/bot_management/domain"
	"bnbot/app/bff/bot_management/infrastructure/external_service/bot/dto"
	"context"
)

func (s *botService) GetAll(ctx context.Context) ([]*domain.Bot, error) {
	bots, err := s.externalbotService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if bots == nil {
		return nil, nil
	}
	botdto := dto.NewEmptyBot()
	res := botdto.FromExternalBotServiceDomainListToDomainList(bots)
	return res, nil
}
