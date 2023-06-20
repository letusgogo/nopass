package config

import (
	_ "embed"
	"fmt"
	"github.com/spf13/viper"
)

//go:embed config-en.yaml
var enConfig string

//go:embed config-zh.yaml
var zhConfig string

type AlgoConfig struct {
	Salt string
}

type ElementConfig struct {
	Name string
	Hint string
}

type Config struct {
	Algo  map[string]*AlgoConfig
	Rules map[string][]*ElementConfig
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return &cfg, nil
}

func GetDefaultConfig(lang string) (string, error) {
	if lang == "zh" {
		return zhConfig, nil
	} else {
		return enConfig, nil
	}

}
