package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigApp struct {
	Loglevel  int    `yaml:"level"`
	Port      string `yaml:"addr" env-default:"8082"`
	Configlog string `yaml:"configlogger"`
}

const cfgPath string = "app/config/config.yaml"

// ReadingConfig reading the config
func ReadingConfig() (*ConfigApp, error) {
	const nm = "[ReadingConfig]"
	var cfg ConfigApp

	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("%s %v", nm, err)
	}

	return &cfg, nil
}
