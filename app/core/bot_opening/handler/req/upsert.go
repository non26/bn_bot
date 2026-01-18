package req

import "bnbot/app/core/bot_opening/domain"

type UpsertRequest struct {
	BotId      string `json:"bot_id"`
	TemplateId string `json:"template_id"`
	ClientId   string `json:"client_id"`
}

func (u *UpsertRequest) ToDomain() *domain.BotOpening {
	return &domain.BotOpening{
		BotId:      u.BotId,
		TemplateId: u.TemplateId,
		ClientId:   u.ClientId,
	}
}
