package req

import "bnbot/app/core/bot/domain"

type DeleteRequest struct {
	BotId string `json:"bot_id"`
}

func (d *DeleteRequest) ToDomain() *domain.Bot {
	return &domain.Bot{
		BotId: d.BotId,
	}
}
