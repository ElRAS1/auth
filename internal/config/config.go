package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPCPort  string `yaml:"grpc_port" env-default:"50051"`
	HTTPPort  string `yaml:"http_port" env-default:"8081"`
	HTTPHost  string `yaml:"http_host" env-default:"localhost"`
	Network   string `yaml:"network" env-default:"tcp"`
	LogLevel  int    `yaml:"level" env-default:"0"`
	ConfigLog string `yaml:"config_logger" env-default:"prod"`
}

const cfgPath string = "config.yaml"

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(cfgPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
