package req

import "bnbot/app/bff/bot_trade_management/domain"

type NewOrderRequest struct {
	TemplateId   string `json:"template_id"`
	ExchangeId   string `json:"exchange_id"`
	BotId        string `json:"bot_id"`
	TradeType    string `json:"trade_type"` // future, spot
	Symbol       string `json:"symbol"`
	Quantity     string `json:"quantity"`
	Side         string `json:"side"`          // buy, sell
	PositionSide string `json:"position_side"` // long, short
	ClientId     string `json:"client_id"`
	AccountId    string `json:"account_id"`
}

func (n *NewOrderRequest) ToDomain() *domain.Trade {
	return &domain.Trade{
		TemplateId:   n.TemplateId,
		ExchangeId:   n.ExchangeId,
		TradeType:    n.TradeType,
		Symbol:       n.Symbol,
		Quantity:     n.Quantity,
		Side:         n.Side,
		PositionSide: n.PositionSide,
		ClientId:     n.ClientId,
		AccountId:    n.AccountId,
	}
}
