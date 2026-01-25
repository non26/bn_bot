package externalservice

import (
	externalbnbotcoreservce "bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core"
	externalbotopensingservice "bnbot/app/core/bot_opening/service"
)

type externalBnBotOpeningService struct {
	service externalbotopensingservice.IService
}

func NewExternalBnBotOpeningService(service externalbotopensingservice.IService) externalbnbotcoreservce.IExternalBnBotOpeningService {
	return &externalBnBotOpeningService{service: service}
}
