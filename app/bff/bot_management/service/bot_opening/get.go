package service

import (
	"bnbot/app/bff/bot_management/domain"
	"context"
	"errors"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *botOpeningService) Get(ctx context.Context, botId string) (*domain.BotOpening, error) {
	botOpening, err := s.externalService.Get(ctx, botId)
	if err != nil {
		return nil, err
	}
	if botOpening == nil {
		return nil, errors.New(appresponse.BOTNOTFOUNDCODE)
	}
	return botOpening, nil
}
