package db_sql_generator

// Config содержит конфигурацию SQL-генератора.
type Config struct {
	Driver string
}

// NewConfig создаёт новую конфигурацию.
func NewConfig() *Config {
	return &Config{}
}

// Configuration загружает конфигурацию из Configurator.
// Если Driver не задан и зарегистрирован только один драйвер, используется он.
func Configuration(config *Config) *Config {
	if config.Driver == "" && aliases.Len() == 1 {
		config.Driver = aliases.Keys()[0]
	}

	return config
}
