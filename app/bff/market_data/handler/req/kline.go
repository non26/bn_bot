package req

import "bnbot/app/bff/market_data/domain"

type KlineRequest struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

func (k *KlineRequest) ToDomain() *domain.Kline {
	return &domain.Kline{
		Symbol:    k.Symbol,
		Interval:  k.Interval,
		StartTime: k.StartTime,
		EndTime:   k.EndTime,
	}
}
