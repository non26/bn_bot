package config

import "github.com/spf13/viper"

type Config struct {
	Env      string         `mapstructure:"environment" json:"environment"`
	Port     int            `mapstructure:"port" json:"port"`
	Id       string         `mapstructure:"id" json:"id"`
	Binance  BinanceService `mapstructure:"binance" json:"binance"`
	DynamoDB DynamoDB       `mapstructure:"dynamodb" json:"dynamodb"`
}

func NewConfig() *Config {
	return &Config{}
}

type BinanceService struct {
	PostionManagement BinancePostionManagement `mapstructure:"postion_management" json:"postion_management"`
	Trade             BinanceTrade             `mapstructure:"trade" json:"trade"`
	MarketData        BinanceMarketData        `mapstructure:"market_data" json:"market_data"`
}

type BinancePostionManagement struct {
	BaseUrl                string `mapstructure:"base_url" json:"base_url"`
	UpdatePositionEndPoint string `mapstructure:"update_position_end_point" json:"update_position_end_point"`
	DeletePositionEndPoint string `mapstructure:"delete_position_end_point" json:"delete_position_end_point"`
	GetPositionEndPoint    string `mapstructure:"get_position_end_point" json:"get_position_end_point"`
}

type BinanceTrade struct {
	BaseUrl          string `mapstructure:"base_url" json:"base_url"`
	NewOrderEndPoint string `mapstructure:"new_order_end_point" json:"new_order_end_point"`
}

type BinanceMarketData struct {
	BaseUrl                  string `mapstructure:"base_url" json:"base_url"`
	GetKlineEndPoint         string `mapstructure:"get_kline_end_point" json:"get_kline_end_point"`
	GetPreviousKlineEndPoint string `mapstructure:"get_previous_kline_end_point" json:"get_previous_kline_end_point"`
}

type DynamoDB struct {
	Region   string `mapstructure:"region" json:"region"`
	Ak       string `mapstructure:"ak" json:"ak"`
	Sk       string `mapstructure:"sk" json:"sk"`
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
}

func LoadConfig() (*Config, error) {
	config := NewConfig()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.Unmarshal(&config)
	return config, nil
}
