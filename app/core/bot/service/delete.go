package service

import "context"

func (s *service) Delete(ctx context.Context, botId string) error {
	return s.repository.Delete(ctx, botId)
}
