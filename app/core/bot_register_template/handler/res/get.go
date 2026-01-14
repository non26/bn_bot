package res

import "bnbot/app/core/bot_register_template/domain"

type GetTemplateResponse struct {
	ExchangeId   string `json:"exchange_id"`
	TemplateId   string `json:"template_id"`
	AllowFutures bool   `json:"allow_futures"`
	AllowSpot    bool   `json:"allow_spot"`
	CreatedAt    string `json:"created_at"`
}

func (g *GetTemplateResponse) FromDomain(template *domain.BotTemplate) {
	g.ExchangeId = template.ExchangeId
	g.TemplateId = template.TemplateId
	g.AllowFutures = template.AllowFutures
	g.AllowSpot = template.AllowSpot
	g.CreatedAt = template.CreatedAt
}
