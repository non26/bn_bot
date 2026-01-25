package externalservice

import (
	externalservice "bnbot/app/bff/bot_management/infrastructure/external_service"
	externalbotservice "bnbot/app/core/bot/service"
)

type botService struct {
	externalbotService externalbotservice.IService
}

func NewBotService(externalbotService externalbotservice.IService) externalservice.IBotManagementService {
	return &botService{externalbotService: externalbotService}
}
