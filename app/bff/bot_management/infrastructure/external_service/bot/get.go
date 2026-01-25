package externalservice

import (
	"bnbot/app/bff/bot_management/domain"
	"bnbot/app/bff/bot_management/infrastructure/external_service/bot/dto"
	"context"
)

func (s *botService) Get(ctx context.Context, botId string) (*domain.Bot, error) {
	bot, err := s.externalbotService.Get(ctx, botId)
	if err != nil {
		return nil, err
	}
	if bot == nil {
		return nil, nil
	}
	botdto := dto.NewEmptyBot()
	res := botdto.FromExternalBotServiceDomainToDomain(bot)
	return res, nil
}
