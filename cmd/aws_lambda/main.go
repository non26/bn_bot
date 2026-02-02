package main

import (
	"bnbot/cmd"
	"bnbot/config"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"bnbot/app/bff"
	"bnbot/app/core"

	bndynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
)

var echoLambda *ginadapter.GinLambda
var _config *config.Config

func init() {

	var err error
	_config, err = cmd.ReadAWSAppLog()
	if err != nil {
		panic(err.Error())
	}

	// dynamodb config
	dynamodbconfig := bndynamodbconfig.NewDynamodbConfig()
	dynamodbendpoint := bndynamodbconfig.NewEndPointResolver(_config.DynamoDB.Region, _config.DynamoDB.Endpoint)
	dynamodbcredential := bndynamodbconfig.NewCredential(_config.DynamoDB.Ak, _config.DynamoDB.Sk)
	var dynamodbclient *dynamodb.Client
	if _config.IsLocal() {
		dynamodbclient = bndynamodbconfig.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewLocal()
	} else {
		dynamodbclient = bndynamodbconfig.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewPrd()
	}

	// echo
	app_gin := gin.Default()
	cmd.HealthCheck(app_gin, _config.HealthCheckMsg)
	// route
	core.RouteCore(app_gin, dynamodbclient)
	bff.BFFRoute(app_gin, dynamodbclient, _config)

	echoLambda = ginadapter.New(app_gin)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
