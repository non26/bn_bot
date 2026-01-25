package externalservice

import (
	"bnbot/app/bff/bot_trade_management/domain"
	"bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bff/trade/req"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *externalBnBffTradeService) NewOrder(ctx context.Context, d_req *domain.Trade) error {
	req := req.NewEmptyNewOrderReq()
	req.FromDomainToExternalServiceReq(d_req)

	jsonReq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s", s.baseUrl, s.newOrderEndpoint)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to new order: %s", response.Status)
	}

	return nil
}
