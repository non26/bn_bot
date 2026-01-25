package externalservice

import (
	"context"
)

func (s *externalBnBotOpeningService) Delete(ctx context.Context, botId string) error {
	opening, err := s.service.Get(ctx, botId)
	if err != nil {
		return err
	}
	if opening == nil {
		return nil
	}
	err = s.service.Delete(ctx, botId)
	if err != nil {
		return err
	}
	return nil
}
