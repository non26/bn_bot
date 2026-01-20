package req

import "bnbot/app/core/bot/domain"

type GetRequest struct {
	BotId string `json:"bot_id"`
}

func (g *GetRequest) ToDomain() *domain.Bot {
	return &domain.Bot{
		BotId: g.BotId,
	}
}
