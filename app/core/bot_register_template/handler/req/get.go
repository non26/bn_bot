package req

import "bnbot/app/core/bot_register_template/domain"

type GetTemplateRequest struct {
	ExchangeId string `json:"exchange_id"`
	TemplateId string `json:"template_id"`
}

func (g *GetTemplateRequest) ToDomain() *domain.BotTemplate {
	return &domain.BotTemplate{
		ExchangeId: g.ExchangeId,
		TemplateId: g.TemplateId,
	}
}
