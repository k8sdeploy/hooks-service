package config

import (
	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/caarlos0/env/v6"
	ConfigBuilder "github.com/keloran/go-config"
)

type Config struct {
	ConfigBuilder.Config
	K8sDeploy
}

func Build() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, bugLog.Error(err)
	}

	c, err := ConfigBuilder.Build(ConfigBuilder.Local, ConfigBuilder.Vault)
	if err != nil {
		return nil, bugLog.Error(err)
	}
	cfg.Config = *c

	if err := BuildK8sDeploy(cfg); err != nil {
		return nil, bugLog.Error(err)
	}

	return cfg, nil
}
