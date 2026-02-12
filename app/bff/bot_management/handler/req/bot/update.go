package req

import "bnbot/app/bff/bot_management/domain"

type UpdateRequest struct {
	BotId      string `json:"bot_id"`
	ExChangeId string `json:"exchange_id"`
	TemplateId string `json:"template_id"`
	IsActive   bool   `json:"is_active"`
	Attributes string `json:"attributes"`
	AccountId  string `json:"account_id"`
}

func (u *UpdateRequest) ToDomain() *domain.Bot {
	return &domain.Bot{
		BotId:      u.BotId,
		ExChangeId: u.ExChangeId,
		TemplateId: u.TemplateId,
		IsActive:   u.IsActive,
		Attributes: u.Attributes,
		AccountId:  u.AccountId,
	}
}
