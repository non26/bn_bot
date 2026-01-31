package route

import (
	"bnbot/app/core/bot/handler"
	bnbot "bnbot/app/core/bot/infrastructure/db/bn_bot"
	"bnbot/app/core/bot/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, client *dynamodb.Client) {
	botRepository := bnbot.NewBnBotRepository(client)
	botService := service.NewService(botRepository)

	group := router.Group("/bn-bot")

	group.POST("/upsert", handler.NewUpsertHandler(botService).Handle)
	group.POST("/delete", handler.NewDeleteHandler(botService).Handle)
	group.POST("/get", handler.NewGetHandler(botService).Handle)
	group.GET("/all", handler.NewGetAllHandler(botService).Handle)
}
