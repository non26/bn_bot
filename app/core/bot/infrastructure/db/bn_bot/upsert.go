package db

import (
	"bnbot/app/core/bot/domain"
	"bnbot/app/core/bot/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
)

func (r *bnBotRepository) Upsert(ctx context.Context, b *domain.Bot) error {
	table := &model.BnBot{}
	table = table.FromDomain(b)

	update_config := dynamodbconfig.NewUpdateTable(table)
	update_config.Set(table.GetExchangeIdField, b.ExChangeId)
	update_config.Set(table.GetTemplateIdField, b.TemplateId)
	update_config.Set(table.GetIsActiveField, b.IsActive)
	update_config.Set(table.GetAttributesField, b.Attributes)
	update_config.Set(table.GetAccountIdField, b.AccountId)
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
