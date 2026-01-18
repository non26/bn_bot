package res

import "bnbot/app/bff/market_data/domain"

type KlineResponse struct {
	Symbol        string `json:"symbol"`
	Interval      string `json:"interval"`
	StartTime     int64  `json:"start_time"`
	EndTime       int64  `json:"end_time"`
	IsGreenCandle bool   `json:"is_green_candle"`
}

func (k *KlineResponse) FromDomain(d *domain.Kline) *KlineResponse {
	if d == nil {
		return nil
	}
	k.Symbol = d.Symbol
	k.Interval = d.Interval
	k.StartTime = d.StartTime
	k.EndTime = d.EndTime
	k.IsGreenCandle = d.IsGreenCandle
	return k
}

func (k *KlineResponse) FromDomainList(d []*domain.Kline) []*KlineResponse {
	if d == nil {
		return nil
	}
	result := []*KlineResponse{}
	for _, item := range d {
		result = append(result, k.FromDomain(item))
	}
	return result
}
