package webserver

// Config ...
type Config struct {
	BindAddr    string `yaml:"bing_addr"`
	LogLevel    string `yaml:"log_level"`
	DatabaseURL string `yaml:"database_url"`
	RedisAddr   string `yaml:"redis_address"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
