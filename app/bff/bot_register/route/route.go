package route

import (
	"bnbot/app/bff/bot_register/handler"
	externalservice "bnbot/app/bff/bot_register/infrastructure/external_service/bot_register_template"
	"bnbot/app/bff/bot_register/service"
	"bnbot/app/core/bot_register_template/infrastructure/db"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, repository db.IBotBNTemplateRepository) {
	_externalservice := externalservice.NewBotRegisterTemplateService(repository)
	_service := service.NewService(_externalservice)

	router.POST("/bot-register/set", handler.NewSetHandler(_service).Handle)
	router.POST("/bot-register/update", handler.NewUpdateHandler(_service).Handle)
	router.POST("/bot-register/delete", handler.NewDeleteHandler(_service).Handle)
	router.GET("/bot-register/get", handler.NewGetHandler(_service).Handle)
	router.GET("/bot-register/get-all", handler.NewGetAllHandler(_service).Handle)
}
