package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPCPort    string `env-default:"50051"     yaml:"grpc_port"`
	HTTPPort    string `env-default:"8081"      yaml:"http_port"`
	HTTPSwagger string `env-default:":8090"     yaml:"http_swagger"`
	HTTPHost    string `env-default:"localhost" yaml:"http_host"`
	Network     string `env-default:"tcp"       yaml:"network"`
	ConfigLog   string `env-default:"prod"      yaml:"config_logger"`
	LogLevel    int    `env-default:"0"         yaml:"level"`
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
