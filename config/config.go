package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AlgoConfig struct {
	Salt string
}

type ElementConfig struct {
	Sort int
	Hint string
}

type Config struct {
	Algo  map[string]*AlgoConfig
	Rules map[string]map[string]ElementConfig
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
