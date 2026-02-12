package externalservice

import (
	externalservice "bnbot/app/bff/bot_management/infrastructure/external_service"
	externalbotcoreservice "bnbot/app/core/bot/service"
)

type botService struct {
	externalbotService externalbotcoreservice.IService
}

func NewBotService(externalbotService externalbotcoreservice.IService) externalservice.IBotExternalCoreService {
	return &botService{externalbotService: externalbotService}
}
