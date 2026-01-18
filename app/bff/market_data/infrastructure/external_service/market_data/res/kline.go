package res

import "bnbot/app/bff/market_data/domain"

type KlineResponse struct {
	Open                  string `json:"open"`
	High                  string `json:"high"`
	Low                   string `json:"low"`
	Close                 string `json:"close"`
	IsCloseHigherThanOpen bool   `json:"is_green_candle"`
}

func (k *KlineResponse) ToDomain() *domain.Kline {
	return &domain.Kline{
		Open:          k.Open,
		High:          k.High,
		Low:           k.Low,
		Close:         k.Close,
		IsGreenCandle: k.IsCloseHigherThanOpen,
	}
}
