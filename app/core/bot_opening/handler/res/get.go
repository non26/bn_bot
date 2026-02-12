package res

import "bnbot/app/core/bot_opening/domain"

type GetResponse struct {
	BotId       string `json:"bot_id"`
	TemplateId  string `json:"template_id"`
	ClientId    string `json:"client_id"`
	BnClientId  string `json:"bn_client_id"`
	Restriction string `json:"restriction"`
}

func (g *GetResponse) FromDomain(b *domain.BotOpening) *GetResponse {
	if b == nil {
		return nil
	}
	g.BotId = b.BotId
	g.TemplateId = b.TemplateId
	g.ClientId = b.ClientId
	g.BnClientId = b.BnClientId
	g.Restriction = b.Restriction
	g.Restriction = b.Restriction
	return g
}
