package db

import (
	"bnbot/app/core/bot_opening/infrastructure/db"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type botOpeningRepository struct {
	client *dynamodb.Client
}

func NewBotOpeningRepository(client *dynamodb.Client) db.IBotOnRunRepository {
	return &botOpeningRepository{client: client}
}
