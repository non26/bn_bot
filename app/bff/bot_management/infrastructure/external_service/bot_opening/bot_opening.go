package externalservice

import (
	externalservice "bnbot/app/bff/bot_management/infrastructure/external_service"
	externalbotopeningcoreservice "bnbot/app/core/bot_opening/service"
)

type externalBotOpeningService struct {
	service externalbotopeningcoreservice.IService
}

func NewExternalBotOpeningService(service externalbotopeningcoreservice.IService) externalservice.IBotOpeningExternalCoreService {
	return &externalBotOpeningService{service: service}
}
