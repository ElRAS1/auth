package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const cfgPath string = "config.yaml"

type Config struct {
	GRPCPort    string `env-default:"50051"     yaml:"grpc_port"`
	HTTPPort    string `env-default:"8081"      yaml:"http_port"`
	SwaggerPort string `env-default:":8090"     yaml:"swagger_port"`
	Host        string `env-default:"localhost" yaml:"host"`
	Network     string `env-default:"tcp"       yaml:"network"`
	ConfigLog   string `env-default:"prod"      yaml:"config_logger"`
	LogLevel    int    `env-default:"0"         yaml:"level"`
}

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(cfgPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
