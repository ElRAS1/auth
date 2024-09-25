package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Server struct {
	LogLevel  int    `yaml:"level" env-default:"0"`
	Port      string `yaml:"addr" env-default:"50051"`
	ConfigLog string `yaml:"configlogger" env-default:"prod"`
	Network   string `yaml:"network" env-default:"tcp"`
}

const cfgPath string = "config.yaml"

// ReadingConfig reading the config
func NewServerCfg() (*Server, error) {
	const nm = "[ReadingConfig]"

	cfg := &Server{}

	err := cleanenv.ReadConfig(cfgPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("%s %v", nm, err)
	}

	return cfg, nil
}
