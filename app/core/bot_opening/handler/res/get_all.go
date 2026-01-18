package res

import "bnbot/app/core/bot_opening/domain"

type GetAllResponse struct {
	Items []*GetResponse `json:"items"`
}

func (g *GetAllResponse) FromDomain(b []*domain.BotOpening) *GetAllResponse {
	if b == nil {
		return nil
	}
	g.Items = []*GetResponse{}
	for _, item := range b {
		_item := &GetResponse{}
		_item.FromDomain(item)
		g.Items = append(g.Items, _item)
	}
	return g
}
