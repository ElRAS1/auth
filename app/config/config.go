package config

import (
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigApp struct {
	Loglevel  int    `yaml:"level"`
	Port      string `yaml:"addr" env-default:"8082"`
	Configlog string `yaml:"configlogger"`
}

const cfgPath string = "../config/config.yaml"

func InitConfig() (*ConfigApp, error) {
	var cfg ConfigApp
	var err error

	workDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cfgPath := filepath.Join(workDir, cfgPath)
	err = cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
