package req

import "bnbot/app/bff/bot_register/domain"

type DeleteRequest struct {
	ExchangeId string `json:"exchange_id"`
	TemplateId string `json:"template_id"`
}

func (d *DeleteRequest) ToDomain() *domain.BotRegister {
	return &domain.BotRegister{
		ExchangeId: d.ExchangeId,
		TemplateId: d.TemplateId,
	}
}
