package req

import "bnbot/app/bff/bot_register/domain"

type UpdateRequest struct {
	ExchangeId   string `json:"exchange_id"`
	TemplateId   string `json:"template_id"`
	AllowFutures bool   `json:"allow_futures"`
	AllowSpot    bool   `json:"allow_spot"`
}

func (u *UpdateRequest) ToDomain() *domain.BotRegister {
	return &domain.BotRegister{
		ExchangeId:   u.ExchangeId,
		TemplateId:   u.TemplateId,
		AllowFutures: u.AllowFutures,
		AllowSpot:    u.AllowSpot,
	}
}
