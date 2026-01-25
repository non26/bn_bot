package externalservice

import (
	"bnbot/app/bff/bot_trade_management/domain"
	"bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core/bot/dto"
	"context"
)

func (s *externalBotService) GetAll(ctx context.Context) ([]*domain.Trade, error) {
	bots, err := s.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if bots == nil {
		return nil, nil
	}
	res := []*domain.Trade{}
	for _, bot := range bots {
		botdto := dto.NewEmptyBot()
		res = append(res, botdto.FromExternalBotServiceDomainToDomain(bot))
	}
	return res, nil
}
