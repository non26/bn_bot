package req

import "bnbot/app/core/bot_opening/domain"

type GetRequest struct {
	BotId string `json:"bot_id"`
}

func (g *GetRequest) ToDomain() *domain.BotOpening {
	return &domain.BotOpening{
		BotId: g.BotId,
	}
}
