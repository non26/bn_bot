package db

import (
	"bnbot/app/core/bot_opening/domain"
	"bnbot/app/core/bot_opening/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *botOpeningRepository) Get(ctx context.Context, botId string) (*domain.BotOpening, error) {
	table := model.BotOpening{}
	table.BotId = botId
	result := model.NewEmptyBotOpening()
	item, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: table.TableName(),
		Key:       table.GetKey(),
	})
	if err != nil {
		return nil, err
	}
	if item.Item == nil {
		return nil, nil
	}
	err = attributevalue.UnmarshalMap(item.Item, result)
	if err != nil {
		return nil, err
	}
	return result.ToDomain(), nil
}
