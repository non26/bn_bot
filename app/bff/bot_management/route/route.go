package route

import (
	bothandler "bnbot/app/bff/bot_management/handler/bot"
	botopeninghandler "bnbot/app/bff/bot_management/handler/bot_opening"
	externalservice "bnbot/app/bff/bot_management/infrastructure/external_service/bot"
	externalbotopeningservice "bnbot/app/bff/bot_management/infrastructure/external_service/bot_opening"
	botservice "bnbot/app/bff/bot_management/service/bot"
	botopeningservice "bnbot/app/bff/bot_management/service/bot_opening"

	externalbotcoreservice "bnbot/app/core/bot/service"
	externalbotopeningcoreservice "bnbot/app/core/bot_opening/service"

	"github.com/gin-gonic/gin"
)

func Route(group *gin.RouterGroup,
	extrenalBotService externalbotcoreservice.IService,
	extrenalBotOpeningCoreService externalbotopeningcoreservice.IService,

) {
	_group := group.Group("/bot-management")

	externalBotService := externalservice.NewBotService(extrenalBotService)
	service := botservice.NewService(externalBotService)
	botGroup := _group.Group("/bot")
	botGroup.POST("/insert", bothandler.NewInsertHandler(service).Handle)
	botGroup.POST("/update", bothandler.NewUpdateHandler(service).Handle)
	botGroup.POST("/delete", bothandler.NewDeleteHandler(service).Handle)
	botGroup.GET("/get", bothandler.NewGetHandler(service).Handle)
	botGroup.GET("/get-all", bothandler.NewGetAllHandler(service).Handle)

	externalBotOpeningService := externalbotopeningservice.NewExternalBotOpeningService(extrenalBotOpeningCoreService)
	botOpeningService := botopeningservice.NewBotOpeningService(externalBotOpeningService)
	botOpeningGroup := _group.Group("/bot-opening")
	botOpeningGroup.POST("/insert", botopeninghandler.NewInsertHandler(botOpeningService).Handle)
	botOpeningGroup.POST("/update", botopeninghandler.NewUpdateHandler(botOpeningService).Handle)
	botOpeningGroup.POST("/delete", botopeninghandler.NewDeleteHandler(botOpeningService).Handle)
	botOpeningGroup.GET("/get", botopeninghandler.NewGetHandler(botOpeningService).Handle)
	botOpeningGroup.GET("/get-all", botopeninghandler.NewGetAllHandler(botOpeningService).Handle)

}
