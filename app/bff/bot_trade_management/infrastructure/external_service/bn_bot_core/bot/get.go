package externalservice

import (
	"bnbot/app/bff/bot_trade_management/domain"
	"bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core/bot/dto"
	"context"
)

func (s *externalBotService) Get(ctx context.Context, botId string) (*domain.Trade, error) {
	bot, err := s.service.Get(ctx, botId)
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
