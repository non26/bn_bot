package service

import (
	"bnbot/app/bff/bot_trade_management/domain"
	"context"
	"errors"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *service) NewOrder(ctx context.Context, d_req *domain.Trade) error {
	isBuyPosition := d_req.IsBuyPosition()
	if isBuyPosition {
		botopening, err := s.botOpeningService.Get(ctx, d_req.BotId)
		if err != nil {
			return err
		}
		if botopening != nil {
			return errors.New(appresponse.BOTFOUNDCODE)
		}
	} else { // sell position
		botopening, err := s.botOpeningService.Get(ctx, d_req.BotId)
		if err != nil {
			return err
		}
		if botopening == nil {
			return errors.New(appresponse.BOTNOTFOUNDCODE)
		}
	}

	err := s.tradeService.NewOrder(ctx, d_req)
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
