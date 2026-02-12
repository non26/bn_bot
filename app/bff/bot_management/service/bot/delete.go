package service

import "context"

func (s *botService) Delete(ctx context.Context, botId string) error {
	err := s.externalService.Delete(ctx, botId)
	if err != nil {
		return err
	}
	return nil
}
