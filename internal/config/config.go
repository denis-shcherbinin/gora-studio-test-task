package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	HttpConfig   *HttpConfig
	SqliteConfig *SqliteConfig
}

func NewConfig(configPath string) (*Config, error) {

	path := strings.Split(configPath, "/")
	viper.AddConfigPath(path[0])

	viper.SetConfigName(path[1])
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New(fmt.Sprintf("error initializing configs: %s", err.Error()))
	}

	return &Config{
		HttpConfig:   httpConfigInstance(),
		SqliteConfig: sqliteConfigInstance(),
	}, nil
}
