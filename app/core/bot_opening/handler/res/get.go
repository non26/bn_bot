package res

import "bnbot/app/core/bot_opening/domain"

type GetResponse struct {
	BotId      string `json:"bot_id"`
	TemplateId string `json:"template_id"`
	ClientId   string `json:"client_id"`
}

func (g *GetResponse) FromDomain(b *domain.BotOpening) *GetResponse {
	if b == nil {
		return nil
	}
	g.BotId = b.BotId
	g.TemplateId = b.TemplateId
	g.ClientId = b.ClientId
	return g
}
