package domain

import (
	"strings"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type Trade struct {
	TemplateId   string
	ExchangeId   string
	BotId        string
	TradeType    string // future, spot
	Symbol       string
	Quantity     string
	Side         string // buy, sell
	PositionSide string // long, short
	ClientId     string
	AccountId    string
	IsBotActve   bool
}

func (t *Trade) IsSpotTrade() bool {
	return strings.ToLower(t.TradeType) == bnconstant.SPOT
}

func (t *Trade) IsBuyPosition() bool {
	if t.IsSpotTrade() {
		return strings.ToLower(t.Side) == bnconstant.BUY
	} else {
		if utils.IsBuyPosition(t.Side, t.PositionSide) {
			return true
		} else {
			return false
		}
	}
}
