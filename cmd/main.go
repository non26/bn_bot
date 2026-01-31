package main

import (
	"bnbot/app/core"
	"bnbot/config"
	"fmt"

	"github.com/gin-gonic/gin"
	bndynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// dynamodb config
	dynamodbconfig := bndynamodbconfig.NewDynamodbConfig()
	dynamodbendpoint := bndynamodbconfig.NewEndPointResolver(config.DynamoDB.Region, config.DynamoDB.Endpoint)
	dynamodbcredential := bndynamodbconfig.NewCredential(config.DynamoDB.Ak, config.DynamoDB.Sk)
	dynamodbclient := bndynamodbconfig.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewLocal()

	router := gin.Default()
	core.RouteCore(router, dynamodbclient)

	router.Run(fmt.Sprintf(":%d", config.Port))
}
