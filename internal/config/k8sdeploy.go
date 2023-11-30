package config

import (
	"github.com/bugfixes/go-bugfixes/logs"
	"github.com/caarlos0/env/v6"
)

type K8sDeploy struct {
	KeyService
}

type KeyService struct {
	Address string `env:"KEY_SERVICE_ADDRESS" envDefault:"key-service.k8sdeploy:8001"`
	Key     string `env:"KEY_SERVICE_KEY" envDefault:""`
}

func BuildK8sDeploy(c *Config) error {
	cfg := &K8sDeploy{}

	if err := env.Parse(cfg); err != nil {
		return logs.Errorf("parse: %v", err)
	}

	c.K8sDeploy = *cfg
	return nil
}
