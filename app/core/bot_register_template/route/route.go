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

	group := router.Group("/bot-bn-template")

	group.POST("/upsert", handler.NewRegisterTemplateHandler(botBNTemplateService).Handle)
	group.POST("/get", handler.NewGetTemplateHandler(botBNTemplateService).Handle)
	group.GET("/all", handler.NewGetAllTemplateHandler(botBNTemplateService).Handle)
	group.POST("/delete", handler.NewDeleteTemplateHandler(botBNTemplateService).Handle)
}
