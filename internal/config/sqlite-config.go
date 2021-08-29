package config

import "github.com/spf13/viper"

const (
	sqliteDriverName     = "sqlite.driver_name"
	sqliteDataSourceName = "sqlite.data_source_name"
)

type SqliteConfig struct {
	DriverName     string
	DataSourceName string
}

func sqliteConfigInstance() *SqliteConfig {
	return &SqliteConfig{
		DriverName:     viper.GetString(sqliteDriverName),
		DataSourceName: viper.GetString(sqliteDataSourceName),
	}
}
