package externalservice

import (
	externalbnbotcoreservce "bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core"
	externalbotservice "bnbot/app/core/bot/service"
)

type externalBotService struct {
	service externalbotservice.IService
}

func NewExternalBotService(service externalbotservice.IService) externalbnbotcoreservce.IExternalBotService {
	return &externalBotService{service: service}
}
