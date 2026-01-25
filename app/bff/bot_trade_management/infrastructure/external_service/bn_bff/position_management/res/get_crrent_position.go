package res

import "bnbot/app/bff/bot_trade_management/domain"

type GetRes struct {
	Symbol       string `json:"symbol"`
	AccountId    string `json:"accountId"`
	PositionSide string `json:"positionSide"`
	Side         string `json:"side"`
	AmountB      string `json:"amountB"`
	ClientId     string `json:"clientId"`
}

func NewEmptyGetRes() *GetRes {
	return &GetRes{}
}

func (g *GetRes) ToDomain() *domain.Trade {
	return &domain.Trade{
		Symbol:       g.Symbol,
		AccountId:    g.AccountId,
		PositionSide: g.PositionSide,
		Side:         g.Side,
		Quantity:     g.AmountB,
		ClientId:     g.ClientId,
	}
}
