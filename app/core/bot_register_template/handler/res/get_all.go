package res

import (
	"bnbot/app/core/bot_register_template/domain"
)

type GetAllTemplateResponse struct {
	Items []*GetTemplateResponse `json:"items"`
}

func (g *GetAllTemplateResponse) FromDomain(templates []*domain.BotTemplate) {
	g.Items = []*GetTemplateResponse{}
	for _, template := range templates {
		item := &GetTemplateResponse{}
		item.FromDomain(template)
		g.Items = append(g.Items, item)
	}
}
