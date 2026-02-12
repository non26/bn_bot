package res

import "bnbot/app/bff/bot_management/domain"

type GetResponse struct {
	BotId      string `json:"bot_id"`
	ExChangeId string `json:"exchange_id"`
	TemplateId string `json:"template_id"`
	IsActive   bool   `json:"is_active"`
	Attributes string `json:"attributes"`
	AccountId  string `json:"account_id"`
}

func (g *GetResponse) FromDomain(b *domain.Bot) *GetResponse {
	if b == nil {
		return nil
	}
	g.BotId = b.BotId
	g.ExChangeId = b.ExChangeId
	g.TemplateId = b.TemplateId
	g.IsActive = b.IsActive
	g.Attributes = b.Attributes
	g.AccountId = b.AccountId
	return g
}
