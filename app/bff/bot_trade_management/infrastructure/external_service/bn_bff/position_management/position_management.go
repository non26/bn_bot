package externalservice

import (
	externalservice "bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bff"
)

type externalBnBffPositionManagementService struct {
	baseUrl                    string
	getCurrentPositionEndpoint string
}

func NewExternalBnBffPositionManagementService(baseUrl string, getCurrentPositionEndpoint string) externalservice.IExternalBnBffPositionManagementService {
	return &externalBnBffPositionManagementService{baseUrl: baseUrl, getCurrentPositionEndpoint: getCurrentPositionEndpoint}
}
