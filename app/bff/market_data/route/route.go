package route

import (
	"bnbot/app/bff/market_data/handler"
	externalservice "bnbot/app/bff/market_data/infrastructure/external_service/market_data"
	"bnbot/app/bff/market_data/service"
	"bnbot/config"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, config *config.Config) {
	marketDataExternalService := externalservice.NewMarketDataService(
		config.Binance.MarketData.BaseUrl,
		config.Binance.MarketData.GetKlineEndPoint,
		config.Binance.MarketData.GetPreviousKlineEndPoint,
	)
	marketDataService := service.NewService(marketDataExternalService)

	router.POST("/kline", handler.NewGetKlineHandler(marketDataService).Handle)
	router.POST("/previous-kline", handler.NewGetPreviousKlineHandler(marketDataService).Handle)
}
