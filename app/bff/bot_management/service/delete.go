package service

import "context"

func (s *service) Delete(ctx context.Context, botId string) error {
	err := s.botService.Delete(ctx, botId)
	if err != nil {
		return err
	}
	return nil
}
