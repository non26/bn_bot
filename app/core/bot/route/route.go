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

	router.POST("/bn-bot/upsert", handler.NewUpsertHandler(botService).Handle)
	router.POST("/bn-bot/delete", handler.NewDeleteHandler(botService).Handle)
	router.GET("/bn-bot/get", handler.NewGetHandler(botService).Handle)
	router.GET("/bn-bot/get-all", handler.NewGetAllHandler(botService).Handle)
}
