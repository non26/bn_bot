package req

import "bnbot/app/bff/bot_management/domain"

type UpsertRequest struct {
	BotId       string `json:"bot_id"`
	TemplateId  string `json:"template_id"`
	ClientId    string `json:"client_id"`
	BnClientId  string `json:"bn_client_id"`
	Restriction string `json:"restriction"`
}

func (u *UpsertRequest) ToDomain() *domain.BotOpening {
	return &domain.BotOpening{
		BotId:       u.BotId,
		TemplateId:  u.TemplateId,
		ClientId:    u.ClientId,
		BnClientId:  u.BnClientId,
		Restriction: u.Restriction,
	}
}
