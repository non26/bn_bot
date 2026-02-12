package service

import (
	"bnbot/app/bff/bot_management/domain"
	"context"
	"errors"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *botOpeningService) Update(ctx context.Context, b *domain.BotOpening) error {
	botOpening, err := s.externalService.Get(ctx, b.BotId)
	if err != nil {
		return err
	}
	if botOpening == nil {
		return errors.New(appresponse.BOTNOTFOUNDCODE)
	}
	return s.externalService.Update(ctx, b)
}
