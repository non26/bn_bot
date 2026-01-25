package externalservice

import (
	"bnbot/app/bff/bot_trade_management/domain"
	"bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core/bn_opening/dto"
	"context"
	"errors"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *externalBnBotOpeningService) Insert(ctx context.Context, b *domain.Trade) error {
	bot, err := s.service.Get(ctx, b.BotId)
	if err != nil {
		return err
	}
	if bot != nil {
		return errors.New(appresponse.BOTFOUNDCODE)
	}

	reqdto := dto.NewEmptyBotOpening()
	reqExternalService := reqdto.FromDomainToExternalBotOpeningServiceDomain(b)
	err = s.service.Upsert(ctx, reqExternalService)
	if err != nil {
		return err
	}
	return nil
}
