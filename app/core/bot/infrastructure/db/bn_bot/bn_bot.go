package db

import (
	"bnbot/app/core/bot/infrastructure/db"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type bnBotRepository struct {
	client *dynamodb.Client
}

func NewBnBotRepository(client *dynamodb.Client) db.IBotRepository {
	return &bnBotRepository{client: client}
}
