package req

import "bnbot/app/bff/bot_register/domain"

type GetRequest struct {
	ExchangeId string `json:"exchange_id"`
	TemplateId string `json:"template_id"`
}

func (g *GetRequest) ToDomain() *domain.BotRegister {
	return &domain.BotRegister{
		ExchangeId: g.ExchangeId,
		TemplateId: g.TemplateId,
	}
}
