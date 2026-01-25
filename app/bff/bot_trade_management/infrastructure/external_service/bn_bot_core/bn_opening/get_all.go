package externalservice

import (
	"bnbot/app/bff/bot_trade_management/domain"
	"bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core/bn_opening/dto"
	"context"
)

func (s *externalBnBotOpeningService) GetAll(ctx context.Context) ([]*domain.Trade, error) {
	openings, err := s.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if openings == nil {
		return nil, nil
	}
	res := []*domain.Trade{}
	for _, opening := range openings {
		openingdto := dto.NewEmptyBotOpening()
		res = append(res, openingdto.FromExternalBotOpeningServiceDomainToDomain(opening))
	}
	return res, nil
}
