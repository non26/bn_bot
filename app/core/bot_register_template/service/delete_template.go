package service

import "context"

func (s *botBNTemplateService) Delete(ctx context.Context, exchangeId string, templateId string) error {
	return s.repository.Delete(ctx, exchangeId, templateId)
}
