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

func (n *NewOrderReq) FromDomainToExternalServiceReq(d *domain.Trade) {

	n.ClientID = d.BnClientID // send bn client id to binance instead of client id for BOT trade
	n.Symbol = d.Symbol
	n.PositionSide = d.PositionSide
	n.Side = d.Side
	n.AmountB = d.Quantity
	n.AccountId = d.AccountId

}
