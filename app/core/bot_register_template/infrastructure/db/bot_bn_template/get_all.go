package db

import (
	"bnbot/app/core/bot_register_template/domain"
	"bnbot/app/core/bot_register_template/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *botBNTemplateRepository) GetAll(ctx context.Context) ([]*domain.BotTemplate, error) {
	result := []*model.BotBNTemplate{}
	table := model.BotBNTemplate{}
	items, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: table.TableName(),
	})
	if err != nil {
		return nil, err
	}
	if len(items.Items) == 0 {
		return nil, nil
	}
	err = attributevalue.UnmarshalListOfMaps(items.Items, &result)
	if err != nil {
		return nil, err
	}
	res := []*domain.BotTemplate{}
	for _, _item := range result {
		res = append(res, _item.ToDomain())
	}
	return res, nil
}
