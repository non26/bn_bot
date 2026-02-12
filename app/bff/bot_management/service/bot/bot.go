package service

import (
	externalservice "bnbot/app/bff/bot_management/infrastructure/external_service"
	service "bnbot/app/bff/bot_management/service"
)

type botService struct {
	externalService externalservice.IBotExternalCoreService
}

func NewService(externalService externalservice.IBotExternalCoreService) service.IBotService {
	return &botService{externalService: externalService}
}
