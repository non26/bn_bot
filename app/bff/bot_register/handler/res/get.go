package res

import "bnbot/app/bff/bot_register/domain"

type GetResponse struct {
	ExchangeId   string `json:"exchange_id"`
	TemplateId   string `json:"template_id"`
	AllowFutures bool   `json:"allow_futures"`
	AllowSpot    bool   `json:"allow_spot"`
	CreatedAt    string `json:"created_at"`
}

func (g *GetResponse) FromDomain(b *domain.BotRegister) *GetResponse {
	if b == nil {
		return nil
	}
	g.ExchangeId = b.ExchangeId
	g.TemplateId = b.TemplateId
	g.AllowFutures = b.AllowFutures
	g.AllowSpot = b.AllowSpot
	g.CreatedAt = b.CreatedAt
	return g
}
