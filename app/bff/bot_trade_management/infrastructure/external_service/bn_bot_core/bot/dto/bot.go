package dto

import (
	"bnbot/app/bff/bot_trade_management/domain"
	externalbotdomain "bnbot/app/core/bot/domain"
)

type Bot struct {
	BotId      string
	ExChangeId string
	TemplateId string
	IsActive   bool
	Attributes string
	AccountId  string
}

func NewEmptyBot() *Bot {
	return &Bot{}
}

func (b *Bot) FromDomainToExternalBotServiceDomain(d *domain.Trade) *externalbotdomain.Bot {
	if d == nil {
		return nil
	}
	return &externalbotdomain.Bot{
		BotId:      d.BotId,
		ExChangeId: d.ExchangeId,
		TemplateId: d.TemplateId,
		IsActive:   d.IsBotActve,
		// Attributes: d.Attributes,
		AccountId: d.AccountId,
	}
}

func (b *Bot) FromExternalBotServiceDomainToDomain(d *externalbotdomain.Bot) *domain.Trade {
	if d == nil {
		return nil
	}
	return &domain.Trade{
		BotId:      d.BotId,
		ExchangeId: d.ExChangeId,
		TemplateId: d.TemplateId,
		IsBotActve: d.IsActive,
		// Attributes: d.Attributes,
		AccountId: d.AccountId,
	}
}
