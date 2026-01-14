package req

import "bnbot/app/core/bot_register_template/domain"

type DeleteTemplateRequest struct {
	ExchangeId string `json:"exchange_id"`
	TemplateId string `json:"template_id"`
}

func (d *DeleteTemplateRequest) ToDomain() *domain.BotTemplate {
	return &domain.BotTemplate{
		ExchangeId: d.ExchangeId,
		TemplateId: d.TemplateId,
	}
}
