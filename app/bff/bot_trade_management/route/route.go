package route

import (
	"bnbot/app/bff/bot_trade_management/handler"
	_externalpositionmanagementservice "bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bff/position_management"
	_externaltradeservice "bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bff/trade"
	_externalbnopeningservice "bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core/bn_opening"
	_externalbotservice "bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core/bot"
	"bnbot/app/bff/bot_trade_management/service"
	externalbotservice "bnbot/app/core/bot/service"
	externalbotopensingservice "bnbot/app/core/bot_opening/service"
	"bnbot/config"

	"github.com/gin-gonic/gin"
)

func Route(
	router *gin.Engine,
	externalBotService externalbotservice.IService,
	externalBotOpeningService externalbotopensingservice.IService,
	config *config.Config,
) {
	bnopeningService := _externalbnopeningservice.NewExternalBnBotOpeningService(externalBotOpeningService)
	botService := _externalbotservice.NewExternalBotService(externalBotService)
	tradeService := _externaltradeservice.NewExternalBnBffTradeService(
		config.Binance.Trade.BaseUrl,
		config.Binance.Trade.NewOrderEndPoint)
	positionManagementService := _externalpositionmanagementservice.NewExternalBnBffPositionManagementService(
		config.Binance.PostionManagement.BaseUrl,
		config.Binance.PostionManagement.GetPositionEndPoint)

	service := service.NewService(
		tradeService,
		positionManagementService,
		bnopeningService,
		botService,
	)

	router.POST("/trade-management/new-order", handler.NewNewOrderHandler(service).Handle)
}
