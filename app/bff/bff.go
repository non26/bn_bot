package bff

import (
	botmanagementroute "bnbot/app/bff/bot_management/route"
	botregistertemplateroute "bnbot/app/bff/bot_register/route"
	bfftraderoute "bnbot/app/bff/bot_trade_management/route"
	externalbotdbcoreservice "bnbot/app/core/bot/infrastructure/db/bn_bot"
	externalbotcoreservice "bnbot/app/core/bot/service"
	externalbotopensingdbcoreservice "bnbot/app/core/bot_opening/infrastructure/db/bot_opening"
	externalbotopensingcoreservice "bnbot/app/core/bot_opening/service"
	externalbotregistertemplatrepocoreservice "bnbot/app/core/bot_register_template/infrastructure/db/bot_bn_template"
	externalbotregistertemplatcoreservice "bnbot/app/core/bot_register_template/service"
	"bnbot/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, dynamoDBClient *dynamodb.Client, config *config.Config) {
	group := router.Group("/bff")

	// bot core service
	externalBotCoreRepository := externalbotdbcoreservice.NewBnBotRepository(dynamoDBClient)
	externalBotCoreService := externalbotcoreservice.NewService(externalBotCoreRepository)

	// bot opening core service
	externalBotOpeningCoreRepository := externalbotopensingdbcoreservice.NewBotOpeningRepository(dynamoDBClient)
	externalBotOpeningCoreService := externalbotopensingcoreservice.NewService(externalBotOpeningCoreRepository)

	// bot register template core service
	externalBotRegisterTemplatCoreRepository := externalbotregistertemplatrepocoreservice.NewBotBNTemplateRepository(dynamoDBClient)
	externalBotRegisterTemplatCoreService := externalbotregistertemplatcoreservice.NewBotBNTemplateService(externalBotRegisterTemplatCoreRepository)

	// bff bot register template
	botregistertemplateroute.Route(group, externalBotRegisterTemplatCoreService)

	// bff bot trade management
	bfftraderoute.Route(group, externalBotCoreService, externalBotOpeningCoreService, config)

	// bff bot management
	botmanagementroute.Route(group, externalBotCoreService, externalBotOpeningCoreService)
}
