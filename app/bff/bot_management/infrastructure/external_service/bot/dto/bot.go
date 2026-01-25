package dto

import (
	"bnbot/app/bff/bot_management/domain"
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

func (b *Bot) ToDomain() *domain.Bot {
	return &domain.Bot{
		BotId:      b.BotId,
		ExChangeId: b.ExChangeId,
		TemplateId: b.TemplateId,
		IsActive:   b.IsActive,
		Attributes: b.Attributes,
		AccountId:  b.AccountId,
	}
}

func (b *Bot) FromDomainToExternalBotServiceDomain(d *domain.Bot) *externalbotdomain.Bot {
	if d == nil {
		return nil
	}
	return &externalbotdomain.Bot{
		BotId:      d.BotId,
		ExChangeId: d.ExChangeId,
		TemplateId: d.TemplateId,
		IsActive:   d.IsActive,
		Attributes: d.Attributes,
		AccountId:  d.AccountId,
	}
}

func (b *Bot) ToExternalServiceDomain() *externalbotdomain.Bot {
	return &externalbotdomain.Bot{
		BotId:      b.BotId,
		ExChangeId: b.ExChangeId,
		TemplateId: b.TemplateId,
		IsActive:   b.IsActive,
		Attributes: b.Attributes,
		AccountId:  b.AccountId,
	}
}

func (b *Bot) FromExternalBotServiceDomainToDomain(d *externalbotdomain.Bot) *domain.Bot {
	if d == nil {
		return nil
	}
	return &domain.Bot{
		BotId:      d.BotId,
		ExChangeId: d.ExChangeId,
		TemplateId: d.TemplateId,
		IsActive:   d.IsActive,
		Attributes: d.Attributes,
		AccountId:  d.AccountId,
	}
}

func (b *Bot) FromExternalBotServiceDomainListToDomainList(d []*externalbotdomain.Bot) []*domain.Bot {
	if len(d) == 0 {
		return nil
	}
	res := []*domain.Bot{}
	for _, item := range d {
		itemdto := NewEmptyBot()
		itemdomain := itemdto.FromExternalBotServiceDomainToDomain(item)
		res = append(res, itemdomain)
	}
	return res
}
