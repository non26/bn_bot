package config

type Config struct {
	Env  string `mapstructure:"environment" json:"environment"`
	Port int    `mapstructure:"port" json:"port"`
	Id   string `mapstructure:"id" json:"id"`
}
