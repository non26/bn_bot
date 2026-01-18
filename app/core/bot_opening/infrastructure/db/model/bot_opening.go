package model

import (
	"bnbot/app/core/bot_opening/domain"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BotOpening struct {
	BotId      string `dynamodbav:"id" dynamodb:"id"`
	TemplateId string `dynamodbav:"template_id" dynamodb:"template_id"`
	ClientId   string `dynamodbav:"client_id" dynamodb:"client_id"`
}

func NewEmptyBotOpening() *BotOpening {
	return &BotOpening{}
}

func (b *BotOpening) TableName() *string {
	return aws.String("bn_bot_opening")
}

func (b *BotOpening) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: b.BotId},
	}
}

func (b *BotOpening) GetBotIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "BotId", "dynamodb")
	return v, t
}

func (b *BotOpening) GetTemplateIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "TemplateId", "dynamodb")
	return v, t
}

func (b *BotOpening) GetClientIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "ClientId", "dynamodb")
	return v, t
}

func (b *BotOpening) ToDomain() *domain.BotOpening {
	return &domain.BotOpening{
		BotId:      b.BotId,
		TemplateId: b.TemplateId,
		ClientId:   b.ClientId,
	}
}
