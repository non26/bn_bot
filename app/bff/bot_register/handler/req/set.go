package req

import "bnbot/app/bff/bot_register/domain"

type SetRequest struct {
	ExchangeId   string `json:"exchange_id"`
	TemplateId   string `json:"template_id"`
	AllowFutures bool   `json:"allow_futures"`
	AllowSpot    bool   `json:"allow_spot"`
}

func (s *SetRequest) ToDomain() *domain.BotRegister {
	return &domain.BotRegister{
		ExchangeId:   s.ExchangeId,
		TemplateId:   s.TemplateId,
		AllowFutures: s.AllowFutures,
		AllowSpot:    s.AllowSpot,
	}
}
