package externalservice

import (
	"bnbot/app/bff/bot_trade_management/domain"
	"bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bff/position_management/req"
	"bnbot/app/bff/bot_trade_management/infrastructure/external_service/bn_bff/position_management/res"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *externalBnBffPositionManagementService) GetCurrentPosition(ctx context.Context, d_req *domain.Trade) (*domain.Trade, error) {

	req := req.NewEmptyGetReq()
	req.FromDomainToExternalServiceReq(d_req)

	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", s.baseUrl, s.getCurrentPositionEndpoint)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get current position: %s", response.Status)
	}

	var res res.GetRes
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res.ToDomain(), nil
}
