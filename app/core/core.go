package core

import (
	botroute "bnbot/app/core/bot/route"
	botopeningroute "bnbot/app/core/bot_opening/route"
	botregistertemplateroute "bnbot/app/core/bot_register_template/route"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, client *dynamodb.Client) {
	group := router.Group("/core")
	botroute.Routes(group, client)
	botopeningroute.Route(group, client)
	botregistertemplateroute.Route(group, client)
}
