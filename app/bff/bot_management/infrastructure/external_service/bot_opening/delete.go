package externalservice

import "context"

func (s *externalBotOpeningService) Delete(ctx context.Context, botId string) error {
	err := s.service.Delete(ctx, botId)
	if err != nil {
		return err
	}
	return nil
}
