package externalservice

import (
	externalservice "bnbot/app/bff/bot_register/infrastructure/external_service"
	"bnbot/app/core/bot_register_template/infrastructure/db"
)

type botRegisterTemplateService struct {
	repository db.IBotBNTemplateRepository
}

func NewBotRegisterTemplateService(repository db.IBotBNTemplateRepository) externalservice.IBotRegisterTemplateService {
	return &botRegisterTemplateService{repository: repository}
}
