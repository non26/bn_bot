package service

import (
	"bnbot/app/bff/bot_management/domain"
	"context"
	"errors"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *service) Insert(ctx context.Context, b *domain.Bot) error {
	bot, err := s.botService.Get(ctx, b.BotId)
	if err != nil {
		return err
	}
	if bot != nil {
		return errors.New(appresponse.BOTFOUNDCODE)
	}
	err = s.botService.Insert(ctx, b)
	if err != nil {
		return err
	}
	return nil
}
