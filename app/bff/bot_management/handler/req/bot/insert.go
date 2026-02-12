package req

import "bnbot/app/bff/bot_management/domain"

type InsertRequest struct {
	BotId      string `json:"bot_id"`
	ExChangeId string `json:"exchange_id"`
	TemplateId string `json:"template_id"`
	IsActive   bool   `json:"is_active"`
	Attributes string `json:"attributes"`
	AccountId  string `json:"account_id"`
}

func (i *InsertRequest) ToDomain() *domain.Bot {
	return &domain.Bot{
		BotId:      i.BotId,
		ExChangeId: i.ExChangeId,
		TemplateId: i.TemplateId,
		IsActive:   i.IsActive,
		Attributes: i.Attributes,
		AccountId:  i.AccountId,
	}
}
