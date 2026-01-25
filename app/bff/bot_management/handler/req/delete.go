package req

import "bnbot/app/bff/bot_management/domain"

type DeleteRequest struct {
	BotId string `json:"bot_id"`
}

func (d *DeleteRequest) ToDomain() *domain.Bot {
	return &domain.Bot{
		BotId: d.BotId,
	}
}
