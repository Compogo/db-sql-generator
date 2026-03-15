package db_sql_generator

type Config struct {
	Driver Driver
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
