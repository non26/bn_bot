package dto

import (
	"bnbot/app/bff/bot_management/domain"
	externalbotopeningdomain "bnbot/app/core/bot_opening/domain"
)

type BotOpeningDTO struct {
	BotId       string
	TemplateId  string
	ClientId    string
	BnClientId  string
	Restriction string
}

func NewEmptyBotOpening() *BotOpeningDTO {
	return &BotOpeningDTO{}
}

func (b *BotOpeningDTO) ToDomain() *domain.BotOpening {
	return &domain.BotOpening{
		BotId:       b.BotId,
		TemplateId:  b.TemplateId,
		ClientId:    b.ClientId,
		BnClientId:  b.BnClientId,
		Restriction: b.Restriction,
	}
}

func (b *BotOpeningDTO) FromExternalBotOpeningServiceDomainToDomain(d *externalbotopeningdomain.BotOpening) *domain.BotOpening {
	return &domain.BotOpening{
		BotId:       d.BotId,
		TemplateId:  d.TemplateId,
		ClientId:    d.ClientId,
		BnClientId:  d.BnClientId,
		Restriction: d.Restriction,
	}
}

func (b *BotOpeningDTO) FromDomainToExternalBotOpeningServiceDomain(d *domain.BotOpening) *externalbotopeningdomain.BotOpening {
	return &externalbotopeningdomain.BotOpening{
		BotId:       d.BotId,
		TemplateId:  d.TemplateId,
		ClientId:    d.ClientId,
		BnClientId:  d.BnClientId,
		Restriction: d.Restriction,
	}
}
