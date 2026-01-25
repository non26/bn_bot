package db

import (
	"bnbot/app/core/bot_register_template/domain"
	"bnbot/app/core/bot_register_template/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
)

func (r *botBNTemplateRepository) Upsert(ctx context.Context, b *domain.BotTemplate) error {
	table := &model.BotBNTemplate{}
	table = table.FromDomain(b)

	update_config := dynamodbconfig.NewUpdateTable(table)
	update_config.Set(table.GetAllowFuturesField, b.AllowFutures)
	update_config.Set(table.GetAllowSpotField, b.AllowSpot)
	expression := update_config.BuildExpression()
	_, err := r.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String("bot_bn_template"),
		Key:                       table.GetKey(),
		UpdateExpression:          expression,
		ExpressionAttributeValues: update_config.GetExpressionAttributeValues(),
	})
	if err != nil {
		return err
	}
	return nil

}
