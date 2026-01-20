package model

import (
	"bnbot/app/core/bot/domain"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BnBot struct {
	Id         string `dynamodbav:"id" dynamodb:"id"`
	ExChangeId string `dynamodbav:"exchange_id" dynamodb:"exchange_id"`
	TemplateId string `dynamodbav:"template_id" dynamodb:"template_id"`
	IsActive   bool   `dynamodbav:"is_active" dynamodb:"is_active"`
	Attributes string `dynamodbav:"attributes" dynamodb:"attributes"`
	AccountId  string `dynamodbav:"account_id" dynamodb:"account_id"`
}

func NewEmptyBnBot() *BnBot {
	return &BnBot{}
}

func (b *BnBot) TableName() *string {
	return aws.String("bn_bot")
}

func (b *BnBot) GetKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: b.Id},
	}
}

func (b *BnBot) GetIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Id", "dynamodb")
	return v, t
}

func (b *BnBot) GetExchangeIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "ExchangeId", "dynamodb")
	return v, t
}

func (b *BnBot) GetTemplateIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "TemplateId", "dynamodb")
	return v, t
}

func (b *BnBot) GetIsActiveField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "IsActive", "dynamodb")
	return v, t
}

func (b *BnBot) GetAttributesField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "Attributes", "dynamodb")
	return v, t
}

func (b *BnBot) GetAccountIdField() (string, reflect.Type) {
	v, t, _ := utils.GetStructTagValueByField(reflect.TypeOf(b).Elem(), "AccountId", "dynamodb")
	return v, t
}

func (b *BnBot) ToDomain() *domain.Bot {
	return &domain.Bot{
		BotId:      b.Id,
		ExChangeId: b.ExChangeId,
		TemplateId: b.TemplateId,
		IsActive:   b.IsActive,
		Attributes: b.Attributes,
		AccountId:  b.AccountId,
	}
}

func (b *BnBot) FromDomain(d *domain.Bot) *BnBot {
	return &BnBot{
		Id:         d.BotId,
		ExChangeId: d.ExChangeId,
		TemplateId: d.TemplateId,
		IsActive:   d.IsActive,
		Attributes: d.Attributes,
		AccountId:  d.AccountId,
	}
}

func (b *BnBot) FromDomainList(d []*domain.Bot) []*BnBot {
	res := []*BnBot{}
	for _, item := range d {
		res = append(res, b.FromDomain(item))
	}
	return res
}

func (b *BnBot) ToDomainList(d []*BnBot) []*domain.Bot {
	res := []*domain.Bot{}
	for _, item := range d {
		res = append(res, item.ToDomain())
	}
	return res
}
