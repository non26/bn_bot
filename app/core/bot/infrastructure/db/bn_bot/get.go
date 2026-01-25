package db

import (
	"bnbot/app/core/bot/domain"
	"bnbot/app/core/bot/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *bnBotRepository) Get(ctx context.Context, botId string) (*domain.Bot, error) {
	table := model.BnBot{}
	table.Id = botId
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
	result := model.NewEmptyBnBot()
	err = attributevalue.UnmarshalMap(item.Item, result)
	if err != nil {
		return nil, err
	}
	return result.ToDomain(), nil
}
