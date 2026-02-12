package service

import (
	"bnbot/app/bff/bot_management/domain"
	"context"
	"errors"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *botOpeningService) Insert(ctx context.Context, b *domain.BotOpening) error {
	botOpening, err := s.externalService.Get(ctx, b.BotId)
	if err != nil {
		return err
	}
	if botOpening != nil {
		return errors.New(appresponse.BOTFOUNDCODE)
	}
	err = s.externalService.Insert(ctx, b)
	if err != nil {
		return err
	}
	return nil
}
