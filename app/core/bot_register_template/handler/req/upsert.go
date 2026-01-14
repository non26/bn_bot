package req

import "bnbot/app/core/bot_register_template/domain"

type UpsertTemplateRequest struct {
	ExchangeId   string `json:"exchange_id"`
	TemplateId   string `json:"template_id"`
	AllowFutures bool   `json:"allow_futures"`
	AllowSpot    bool   `json:"allow_spot"`
}

func (u *UpsertTemplateRequest) ToDomain() *domain.BotTemplate {
	return &domain.BotTemplate{
		ExchangeId:   u.ExchangeId,
		TemplateId:   u.TemplateId,
		AllowFutures: u.AllowFutures,
		AllowSpot:    u.AllowSpot,
	}
}
