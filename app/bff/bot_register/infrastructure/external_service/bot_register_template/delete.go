package externalservice

import (
	"context"
)

func (s *botRegisterTemplateService) Delete(ctx context.Context, exchangeId string, templateId string) error {
	err := s.repository.Delete(ctx, exchangeId, templateId)
	if err != nil {
		return err
	}

	return nil
}
