package model

import (
	"bnbot/app/core/bot_register_template/domain"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BotBNTemplate struct {
	ExchangeId   string `dynamodbav:"exchange_id" dynamodb:"exchange_id"`
	TemplateId   string `dynamodbav:"template_id" dynamodb:"template_id"`
	AllowFutures bool   `dynamodbav:"allow_futures" dynamodb:"allow_futures"`
	AllowSpot    bool   `dynamodbav:"allow_spot" dynamodb:"allow_spot"`
	CreatedAt    string `dynamodbav:"created_at" dynamodb:"created_at"`
}

func NewEmptyBotBNTemplate() *BotBNTemplate {
	return &BotBNTemplate{}
}

func (b *BotBNTemplate) TableName() string {
	return "bot_bn_template"
}

func (b *BotBNTemplate) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"exchange_id": &types.AttributeValueMemberS{Value: b.ExchangeId},
		"template_id": &types.AttributeValueMemberS{Value: b.TemplateId},
	}
}

func (b *BotBNTemplate) GetExchangeIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "ExchangeId", "dynamodb")
	return v, t
}

func (b *BotBNTemplate) GetTemplateIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "TemplateId", "dynamodb")
	return v, t
}

func (b *BotBNTemplate) GetAllowFuturesField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AllowFutures", "dynamodb")
	return v, t
}

func (b *BotBNTemplate) GetAllowSpotField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AllowSpot", "dynamodb")
	return v, t
}

func (b *BotBNTemplate) GetCreatedAtField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "CreatedAt", "dynamodb")
	return v, t
}

func (b *BotBNTemplate) FromDomain(d *domain.BotTemplate) *BotBNTemplate {
	return &BotBNTemplate{
		ExchangeId:   d.ExchangeId,
		TemplateId:   d.TemplateId,
		AllowFutures: d.AllowFutures,
		AllowSpot:    d.AllowSpot,
		CreatedAt:    d.CreatedAt,
	}
}

func (b *BotBNTemplate) ToDomain() *domain.BotTemplate {
	if b == nil {
		return nil
	}
	return &domain.BotTemplate{
		ExchangeId:   b.ExchangeId,
		TemplateId:   b.TemplateId,
		AllowFutures: b.AllowFutures,
		AllowSpot:    b.AllowSpot,
		CreatedAt:    b.CreatedAt,
	}
}
