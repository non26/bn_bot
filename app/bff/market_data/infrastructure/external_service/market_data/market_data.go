package externalservice

import externalservice "bnbot/app/bff/market_data/infrastructure/external_service"

type marketDataService struct {
	baseUrl                  string
	getKlineEndPoint         string
	getPreviousKlineEndPoint string
}

func NewMarketDataService(
	baseUrl string,
	getKlineEndPoint string,
	getPreviousKlineEndPoint string) externalservice.IMarketDataService {
	return &marketDataService{
		baseUrl:                  baseUrl,
		getKlineEndPoint:         getKlineEndPoint,
		getPreviousKlineEndPoint: getPreviousKlineEndPoint,
	}
}
