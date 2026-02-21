package service

import (
	"bnbot/app/bff/bot_trade_management/domain"
	"context"
	"errors"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *service) NewOrder(ctx context.Context, d_req *domain.Trade) error {
	isBuyPosition := d_req.IsBuyPosition()
	var botopening *domain.Trade
	var err error
	if isBuyPosition {
		botopening, err = s.botOpeningService.Get(ctx, d_req.BotId)
		if err != nil {
			return err
		}
		if botopening != nil {
			return errors.New(appresponse.BOTFOUNDCODE)
		}
	} else { // sell position
		botopening, err = s.botOpeningService.Get(ctx, d_req.BotId)
		if err != nil {
			return err
		}
		if botopening == nil {
			return errors.New(appresponse.BOTNOTFOUNDCODE)
		}
		if botopening.ClientId != d_req.ClientId {
			return errors.New(appresponse.NotFoundOpeningPositionErrorCode)
		}

		// if botopening.TemplateId != d_req.TemplateId {
		// 	return errors.New(appresponse.BOTNOTREGISTEREDCODE)
		// }

		if botopening.BotId != d_req.BotId {
			return errors.New(appresponse.BOTNOTFOUNDCODE)
		}
	}

	err = s.tradeService.NewOrder(ctx, d_req)
	if err != nil {
		return err
	}

	if isBuyPosition {
		err = s.botOpeningService.Insert(ctx, d_req)
		if err != nil {
			return err
		}
	} else {
		err = s.botOpeningService.Delete(ctx, d_req.BotId)
		if err != nil {
			return err
		}
	}

	return nil
}
