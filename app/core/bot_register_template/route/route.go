package route

import (
	"bnbot/app/core/bot_register_template/handler"
	botbntemplate "bnbot/app/core/bot_register_template/infrastructure/db/bot_bn_template"
	"bnbot/app/core/bot_register_template/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, client *dynamodb.Client) {

	botBNTemplateRepository := botbntemplate.NewBotBNTemplateRepository(client)

	botBNTemplateService := service.NewBotBNTemplateService(botBNTemplateRepository)

	router.POST("/bot-bn-template/upsert", handler.NewRegisterTemplateHandler(botBNTemplateService).Handle)
	router.POST("/bot-bn-template", handler.NewGetTemplateHandler(botBNTemplateService).Handle)
	router.GET("/bot-bn-template/all", handler.NewGetAllTemplateHandler(botBNTemplateService).Handle)
	router.POST("/bot-bn-template/delete", handler.NewDeleteTemplateHandler(botBNTemplateService).Handle)
}
