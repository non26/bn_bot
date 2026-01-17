package service

import "context"

func (s *service) Delete(ctx context.Context, exchangeId string, templateId string) error {
	err := s.botRegisterTemplateService.Delete(ctx, exchangeId, templateId)
	if err != nil {
		return err
	}

	return nil
}
