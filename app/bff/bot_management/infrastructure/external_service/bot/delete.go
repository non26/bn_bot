package externalservice

import "context"

func (s *botService) Delete(ctx context.Context, botId string) error {
	err := s.externalbotService.Delete(ctx, botId)
	if err != nil {
		return err
	}
	return nil
}
