package bnbot

import (
	"bnbot/app/core/bot/domain"
	"bnbot/app/core/bot/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *bnBotRepository) GetAll(ctx context.Context) ([]*domain.Bot, error) {
	table := model.BnBot{}
	items, err := r.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: table.TableName(),
	})
	if err != nil {
		return nil, err
	}
	if len(items.Items) == 0 {
		return nil, nil
	}
	result := []*model.BnBot{}
	err = attributevalue.UnmarshalListOfMaps(items.Items, result)
	if err != nil {
		return nil, err
	}
	res := model.NewEmptyBnBot()
	return res.ToDomainList(result), nil
}
