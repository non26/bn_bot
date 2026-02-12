package service

import (
	externalservice "bnbot/app/bff/bot_management/infrastructure/external_service"
	service "bnbot/app/bff/bot_management/service"
)

type botOpeningService struct {
	externalService externalservice.IBotOpeningExternalCoreService
}

func NewBotOpeningService(externalService externalservice.IBotOpeningExternalCoreService) service.IBotOpeningService {
	return &botOpeningService{externalService: externalService}
}
