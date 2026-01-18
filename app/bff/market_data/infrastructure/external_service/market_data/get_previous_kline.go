package externalservice

import (
	"bnbot/app/bff/market_data/domain"
	"bnbot/app/bff/market_data/infrastructure/external_service/market_data/req"
	"bnbot/app/bff/market_data/infrastructure/external_service/market_data/res"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (m *marketDataService) GetPreviousKline(ctx context.Context, d *domain.Kline) (*domain.Kline, error) {
	url := fmt.Sprintf("%s%s", m.baseUrl, m.getPreviousKlineEndPoint)
	reqKline := req.KlineRequest{}
	reqKline.FromDomain(d)
	body, err := json.Marshal(reqKline)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
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
		return nil, fmt.Errorf("failed to get previous kline: %s", response.Status)
	}

	var resKline *res.KlineResponse
	err = json.NewDecoder(response.Body).Decode(&resKline)
	if err != nil {
		return nil, err
	}

	return resKline.ToDomain(), nil
}
