package service

import (
	"context"
	"errors"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

func (s *botOpeningService) Delete(ctx context.Context, botId string) error {
	botOpening, err := s.externalService.Get(ctx, botId)
	if err != nil {
		return err
	}
	if botOpening == nil {
		return errors.New(appresponse.BOTNOTFOUNDCODE)
	}
	err = s.externalService.Delete(ctx, botId)
	if err != nil {
		return err
	}
	return nil
}
