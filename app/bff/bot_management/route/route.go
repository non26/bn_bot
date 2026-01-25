package route

import (
	"bnbot/app/bff/bot_management/handler"
	externalservice "bnbot/app/bff/bot_management/infrastructure/external_service/bot"
	"bnbot/app/bff/bot_management/service"
	externalbotservice "bnbot/app/core/bot/service"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, extrenalBotService externalbotservice.IService) {
	externalBotService := externalservice.NewBotService(extrenalBotService)
	service := service.NewService(externalBotService)

	router.POST("/bot-management/insert", handler.NewInsertHandler(service).Handle)
	router.POST("/bot-management/update", handler.NewUpdateHandler(service).Handle)
	router.POST("/bot-management/delete", handler.NewDeleteHandler(service).Handle)
	router.GET("/bot-management/get", handler.NewGetHandler(service).Handle)
	router.GET("/bot-management/get-all", handler.NewGetAllHandler(service).Handle)

}
