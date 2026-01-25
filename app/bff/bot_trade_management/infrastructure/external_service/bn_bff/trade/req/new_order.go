package req

import "bnbot/app/bff/bot_trade_management/domain"

type NewOrderReq struct {
	ClientID     string `json:"client_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
	Side         string `json:"side"`
	AmountB      string `json:"amount_b"`
	AccountId    string `json:"account_id"`
}

func NewEmptyNewOrderReq() *NewOrderReq {
	return &NewOrderReq{}
}

func (n *NewOrderReq) FromDomainToExternalServiceReq(d *domain.Trade) *NewOrderReq {
	return &NewOrderReq{
		ClientID:     d.ClientId,
		Symbol:       d.Symbol,
		PositionSide: d.PositionSide,
		Side:         d.Side,
		AmountB:      d.Quantity,
		AccountId:    d.AccountId,
	}
}
