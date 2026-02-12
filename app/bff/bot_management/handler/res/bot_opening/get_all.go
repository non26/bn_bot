package res

import "bnbot/app/bff/bot_management/domain"

type GetAllResponse struct {
	Items []*GetResponse `json:"items"`
}

func NewGetAllResponse() *GetAllResponse {
	return &GetAllResponse{
		Items: []*GetResponse{},
	}
}

func (g *GetAllResponse) FromDomainList(b []*domain.BotOpening) *GetAllResponse {
	if len(b) == 0 {
		return nil
	}
	for _, item := range b {
		_item := &GetResponse{}
		_item.FromDomain(item)
		g.Items = append(g.Items, _item)
	}

	return g
}
