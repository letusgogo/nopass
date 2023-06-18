package config

import (
	"github.com/letusgogo/nopass/algo"
	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func New() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return &Config{v: v}
}

type Rule struct {
	Name     string
	Elements map[string]string
}

func (c *Config) GetRules() []Rule {
	return nil
}

func (c *Config) GetAlgo(name string) algo.Algorithm {
	return nil
}
