package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		App   `yaml:"app"`
		MySQL `yaml:"mysql"`
		Host  `yaml:"host"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"`
		Version string `env-requred:"true" yaml:"version"`
	}

	MySQL struct {
		User   string `env-required:"true" yaml:"user"`
		Passwd string `env-required:"true" yaml:"passwd"`
		Net    string `env-required:"true" yaml:"net"`
		Addr   string `env-required:"true" yaml:"addr"`
		DBName string `env-required:"true" yaml:"dbname"`
	}

	Host struct {
		Port string `env-required:"true" yaml:"port"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	path := "./config/local.yaml"

	err := cleanenv.ReadConfig(path, cfg)

	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
