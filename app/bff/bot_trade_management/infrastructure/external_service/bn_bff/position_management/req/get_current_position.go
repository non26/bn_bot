package req

import "bnbot/app/bff/bot_trade_management/domain"

type GetReq struct {
	Symbol       string `json:"symbol"`
	AccountId    string `json:"accountId"`
	PositionSide string `json:"positionSide"`
}

func NewEmptyGetReq() *GetReq {
	return &GetReq{}
}

func (g *GetReq) ToDomain() *domain.Trade {
	return &domain.Trade{
		Symbol:       g.Symbol,
		AccountId:    g.AccountId,
		PositionSide: g.PositionSide,
	}
}

func (g *GetReq) FromDomainToExternalServiceReq(d *domain.Trade) *GetReq {
	return &GetReq{
		Symbol:       d.Symbol,
		AccountId:    d.AccountId,
		PositionSide: d.PositionSide,
	}
}
