package botbntemplate

import (
	"bnbot/app/core/bot_register_template/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *botBNTemplateRepository) Delete(ctx context.Context, exchangeId string, templateId string) error {
	table := model.BotBNTemplate{}
	table.ExchangeId = exchangeId
	table.TemplateId = templateId
	_, err := r.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String("bot_bn_template"),
		Key:       table.GetKey(),
	})
	if err != nil {
		return err
	}
	return err
}
