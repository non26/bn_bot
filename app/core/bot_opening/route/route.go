package route

import (
	"bnbot/app/core/bot_opening/handler"
	db "bnbot/app/core/bot_opening/infrastructure/db/bot_opening"
	"bnbot/app/core/bot_opening/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, client *dynamodb.Client) {
	botOpeningRepository := db.NewBotOpeningRepository(client)
	botOpeningService := service.NewService(botOpeningRepository)

	router.POST("/bot-opening/upsert", handler.NewUpsertHandler(botOpeningService).Handle)
	router.POST("/bot-opening/delete", handler.NewDeleteHandler(botOpeningService).Handle)
	router.GET("/bot-opening/get", handler.NewGetHandler(botOpeningService).Handle)
	router.GET("/bot-opening/get-all", handler.NewGetAllHandler(botOpeningService).Handle)
}
