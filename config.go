package db_sql_generator

import "github.com/Compogo/db-client/driver"

type Config struct {
	Driver driver.Driver
}

func NewConfig() *Config {
	return &Config{}
}

func Configuration(config *Config) *Config {
	if config.Driver == "" && aliases.Len() == 1 {
		config.Driver = aliases.Keys()[0]
	}

	return config
}
