package req

import "bnbot/app/bff/bot_management/domain"

type DeleteRequest struct {
	BotId string `json:"bot_id"`
}

func (d *DeleteRequest) ToDomain() *domain.BotOpening {
	return &domain.BotOpening{
		BotId: d.BotId,
	}
}
