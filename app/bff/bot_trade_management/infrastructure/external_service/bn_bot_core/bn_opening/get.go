package externalservice

import (
	"bnbot/app/bff/bot_trade_management/domain"
	"bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core/bn_opening/dto"
	"context"
)

func (s *externalBnBotOpeningService) Get(ctx context.Context, botId string) (*domain.Trade, error) {
	opening, err := s.service.Get(ctx, botId)
	if err != nil {
		return nil, err
	}
	if opening == nil {
		return nil, nil
	}
	openingdto := dto.NewEmptyBotOpening()
	res := openingdto.FromExternalBotOpeningServiceDomainToDomain(opening)
	return res, nil
}
