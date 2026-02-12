package route

import (
	"bnbot/app/core/bot_opening/handler"
	db "bnbot/app/core/bot_opening/infrastructure/db/bot_opening"
	"bnbot/app/core/bot_opening/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Route(group *gin.RouterGroup, client *dynamodb.Client) {
	botOpeningRepository := db.NewBotOpeningRepository(client)
	botOpeningService := service.NewService(botOpeningRepository)

	group = group.Group("/bn-opening")

	group.POST("/upsert", handler.NewUpsertHandler(botOpeningService).Handle)
	group.POST("/delete", handler.NewDeleteHandler(botOpeningService).Handle)
	group.POST("/get", handler.NewGetHandler(botOpeningService).Handle)
	group.GET("/all", handler.NewGetAllHandler(botOpeningService).Handle)
}
