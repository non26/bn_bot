package route

import (
	"bnbot/app/bff/bot_register/handler"
	externalservice "bnbot/app/bff/bot_register/infrastructure/external_service/bot_register_template"
	"bnbot/app/bff/bot_register/service"
	externalbotregistertemplatcoreservice "bnbot/app/core/bot_register_template/service"

	"github.com/gin-gonic/gin"
)

func Route(group *gin.RouterGroup, externalBotRegisterTemplatCoreService externalbotregistertemplatcoreservice.IBotBNTemplateService) {
	_externalservice := externalservice.NewBotRegisterTemplateService(externalBotRegisterTemplatCoreService)
	_service := service.NewService(_externalservice)

	group.POST("/bot-register/set", handler.NewSetHandler(_service).Handle)
	group.POST("/bot-register/update", handler.NewUpdateHandler(_service).Handle)
	group.POST("/bot-register/delete", handler.NewDeleteHandler(_service).Handle)
	group.GET("/bot-register/get", handler.NewGetHandler(_service).Handle)
	group.GET("/bot-register/get-all", handler.NewGetAllHandler(_service).Handle)
}
