package db

import (
	"bnbot/app/core/bot_register_template/domain"
	"bnbot/app/core/bot_register_template/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *botBNTemplateRepository) Get(ctx context.Context, exchangeId string, templateId string) (*domain.BotTemplate, error) {
	table := model.BotBNTemplate{}
	table.ExchangeId = exchangeId
	table.TemplateId = templateId
	result := model.NewEmptyBotBNTemplate()
	item, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("bot_bn_template"),
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
	return table.ToDomain(), nil
}
