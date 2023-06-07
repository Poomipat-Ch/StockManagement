package configs

import (
	"strings"

	"github.com/Poomipat-Ch/StockManagement/pkg/postgres"
	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     uint64
	Postgres *postgres.Config
}

func InitConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
