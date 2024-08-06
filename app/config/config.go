package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigApp struct {
	Loglevel  int    `yaml:"level"`
	Port      string `yaml:"addr" env-default:"8082"`
	Configlog string `yaml:"configlogger"`
}

const cfgPath string = "app/config/config.yaml"

func InitConfig() (*ConfigApp, error) {
	var cfg ConfigApp
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
