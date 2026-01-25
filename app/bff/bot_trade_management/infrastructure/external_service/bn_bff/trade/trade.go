package externalservice

import (
	externalservice "bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bff"
)

type externalBnBffTradeService struct {
	baseUrl          string
	newOrderEndpoint string
}

func NewExternalBnBffTradeService(
	baseUrl string,
	newOrderEndpoint string,
) externalservice.IExternalBnBffTradeService {
	return &externalBnBffTradeService{
		baseUrl:          baseUrl,
		newOrderEndpoint: newOrderEndpoint,
	}
}
