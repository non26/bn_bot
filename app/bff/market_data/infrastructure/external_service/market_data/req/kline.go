package req

import "bnbot/app/bff/market_data/domain"

type KlineRequest struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

func (k *KlineRequest) FromDomain(d *domain.Kline) *KlineRequest {
	return &KlineRequest{
		Symbol:    d.Symbol,
		Interval:  d.Interval,
		StartTime: d.StartTime,
		EndTime:   d.EndTime,
	}
}
