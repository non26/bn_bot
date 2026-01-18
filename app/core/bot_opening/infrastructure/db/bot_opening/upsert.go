package db

import (
	"bnbot/app/core/bot_opening/domain"
	"bnbot/app/core/bot_opening/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
)

func (r *botOpeningRepository) Upsert(ctx context.Context, b *domain.BotOpening) error {
	table := &model.BotOpening{}
	table.BotId = b.BotId
	table.TemplateId = b.TemplateId
	table.ClientId = b.ClientId

	update_config := dynamodbconfig.NewUpdateTable(table)
	update_config.Set(table.GetTemplateIdField, b.TemplateId)
	update_config.Set(table.GetClientIdField, b.ClientId)
	expression := update_config.BuildExpression()
	_, err := r.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 table.TableName(),
		Key:                       table.GetKey(),
		UpdateExpression:          expression,
		ExpressionAttributeValues: update_config.GetExpressionAttributeValues(),
	})
	if err != nil {
		return err
	}
	return nil
}
