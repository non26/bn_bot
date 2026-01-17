package dto

import (
	"bnbot/app/bff/bot_register/domain"
	externaldomain "bnbot/app/core/bot_register_template/domain"
)

type BotRegisterTemplateDTO struct {
	ExchangeId   string
	TemplateId   string
	AllowFutures bool
	AllowSpot    bool
}

func (b *BotRegisterTemplateDTO) ToDomain() *domain.BotRegister {
	return &domain.BotRegister{
		ExchangeId:   b.ExchangeId,
		TemplateId:   b.TemplateId,
		AllowFutures: b.AllowFutures,
		AllowSpot:    b.AllowSpot,
	}
}

func (b *BotRegisterTemplateDTO) FromDomain(d *domain.BotRegister) *BotRegisterTemplateDTO {
	return &BotRegisterTemplateDTO{
		ExchangeId:   d.ExchangeId,
		TemplateId:   d.TemplateId,
		AllowFutures: d.AllowFutures,
		AllowSpot:    d.AllowSpot,
	}
}

func (b *BotRegisterTemplateDTO) FromExternalServiceDomain(d *externaldomain.BotTemplate) *BotRegisterTemplateDTO {
	return &BotRegisterTemplateDTO{
		ExchangeId:   d.ExchangeId,
		TemplateId:   d.TemplateId,
		AllowFutures: d.AllowFutures,
		AllowSpot:    d.AllowSpot,
	}
}

func (b *BotRegisterTemplateDTO) ToExternalServiceDomain() *externaldomain.BotTemplate {
	return &externaldomain.BotTemplate{
		ExchangeId:   b.ExchangeId,
		TemplateId:   b.TemplateId,
		AllowFutures: b.AllowFutures,
		AllowSpot:    b.AllowSpot,
	}
}
