package util

import "github.com/spf13/viper"

type Config struct {
	DbDriver     string `mapstructure:"DATA_BASE_DRIVER"`
	DbSource     string `mapstructure:"DATA_BASE_SOURCE"`
	SeverAddress string `mapstructure:"SERVER_URL"`
	MongoDbUri   string `mapstructure:"MONGO_DB_URI"`
	RedisUrl     string `mapstructure:"REDIS_URL"`
	RedisPass    string `mapstructure:"REDIS_PASSWORD"`
	RedisDb      int    `mapstructure:"REDIS_DB"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
