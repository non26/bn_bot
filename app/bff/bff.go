package bff

import (
	bfftraderoute "bnbot/app/bff/bot_trade_management/route"
	externalbotdbcoreservice "bnbot/app/core/bot/infrastructure/db/bn_bot"
	externalbotcoreservice "bnbot/app/core/bot/service"
	externalbotopensingdbcoreservice "bnbot/app/core/bot_opening/infrastructure/db/bot_opening"
	externalbotopensingcoreservice "bnbot/app/core/bot_opening/service"
	"bnbot/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func BFFRoute(router *gin.Engine, dynamoDBClient *dynamodb.Client, config *config.Config) {
	group := router.Group("/bff")

	externalBotCoreRepository := externalbotdbcoreservice.NewBnBotRepository(dynamoDBClient)
	externalBotCoreService := externalbotcoreservice.NewService(externalBotCoreRepository)

	externalBotOpeningRepository := externalbotopensingdbcoreservice.NewBotOpeningRepository(dynamoDBClient)
	externalBotOpeningService := externalbotopensingcoreservice.NewService(externalBotOpeningRepository)

	bfftraderoute.Route(group, externalBotCoreService, externalBotOpeningService, config)
}
