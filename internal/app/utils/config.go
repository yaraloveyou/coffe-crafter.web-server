package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RefreshExp int64  `yaml:"refresh_exp"`
	AccessExp  int64  `yaml:"access_exp"`
	SecretKey  string `yaml:"secret_key"`
}

var (
	config *Config
)

func LoadConfig(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	return nil
}
