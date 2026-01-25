package db

import (
	"bnbot/app/core/bot/infrastructure/db/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *bnBotRepository) Delete(ctx context.Context, botId string) error {
	table := model.BnBot{}
	table.Id = botId
	_, err := r.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: table.TableName(),
		Key:       table.GetKey(),
	})
	if err != nil {
		return err
	}
	return nil
}
