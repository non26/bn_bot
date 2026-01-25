package dto

import (
	"bnbot/app/bff/bot_trade_management/domain"
	externalbotopeningdomain "bnbot/app/core/bot_opening/domain"
)

type BotOpening struct {
	BotId      string
	TemplateId string
	ClientId   string
}

func NewEmptyBotOpening() *BotOpening {
	return &BotOpening{}
}

func (b *BotOpening) FromExternalBotOpeningServiceDomainToDomain(d *externalbotopeningdomain.BotOpening) *domain.Trade {
	if d == nil {
		return nil
	}
	return &domain.Trade{
		BotId:      d.BotId,
		TemplateId: d.TemplateId,
		ClientId:   d.ClientId,
	}
}

func (b *BotOpening) FromDomainToExternalBotOpeningServiceDomain(d *domain.Trade) *externalbotopeningdomain.BotOpening {
	if d == nil {
		return nil
	}
	return &externalbotopeningdomain.BotOpening{
		BotId:      d.BotId,
		TemplateId: d.TemplateId,
		ClientId:   d.ClientId,
	}
}
