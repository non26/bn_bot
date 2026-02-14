package service

import (
	"bnbot/app/bff/bot_trade_management/domain"
	externalbnffservice "bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bff"
	externalbnbotcoreservice "bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bot_core"
	"context"
)

type IService interface {
	NewOrder(ctx context.Context, d_req *domain.Trade) error
}

type service struct {
	tradeService      externalbnffservice.IExternalBnBffTradeService
	botOpeningService externalbnbotcoreservice.IExternalBnBotOpeningService
	botService        externalbnbotcoreservice.IExternalBotService
}

func NewService(
	tradeService externalbnffservice.IExternalBnBffTradeService,
	botOpeningService externalbnbotcoreservice.IExternalBnBotOpeningService,
	botService externalbnbotcoreservice.IExternalBotService,
) IService {
	return &service{
		tradeService:      tradeService,
		botOpeningService: botOpeningService,
		botService:        botService,
	}
}
