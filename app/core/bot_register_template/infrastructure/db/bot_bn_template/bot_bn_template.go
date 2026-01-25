package db

import (
	"bnbot/app/core/bot_register_template/infrastructure/db"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type botBNTemplateRepository struct {
	client *dynamodb.Client
}

func NewBotBNTemplateRepository(client *dynamodb.Client) db.IBotBNTemplateRepository {
	return &botBNTemplateRepository{client: client}
}
