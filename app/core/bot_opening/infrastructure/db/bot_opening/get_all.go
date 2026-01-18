package db

import (
	"bnbot/app/core/bot_opening/domain"
	"bnbot/app/core/bot_opening/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *botOpeningRepository) GetAll(ctx context.Context) ([]*domain.BotOpening, error) {
	table := model.BotOpening{}
	items, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: table.TableName(),
	})
	if err != nil {
		return nil, err
	}
	if len(items.Items) == 0 {
		return nil, nil
	}
	result := []*model.BotOpening{}
	err = attributevalue.UnmarshalListOfMaps(items.Items, result)
	if err != nil {
		return nil, err
	}
	res := []*domain.BotOpening{}
	for _, item := range result {
		res = append(res, item.ToDomain())
	}
	return res, nil
}
