package req

import "bnbot/app/bff/bot_management/domain"

type GetRequest struct {
	BotId string `json:"bot_id"`
}

func (g *GetRequest) ToDomain() *domain.BotOpening {
	return &domain.BotOpening{
		BotId: g.BotId,
	}
}
