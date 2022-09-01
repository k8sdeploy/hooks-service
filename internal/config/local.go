package config

import "github.com/caarlos0/env/v6"

type OrchestratorService struct {
	Key     string `env:"HOOKS_KEY"`
	Secret  string `env:"HOOKS_SECRET"`
	Address string `env:"ORCHESTRATOR_ADDRESS" envDefault:"orchestrator.k8sdeploy:8001"`
}
type KeyService struct {
	Key     string `env:"HOOKS_KEY"`
	Secret  string `env:"HOOKS_SECRET"`
	Address string `env:"KEY_SERVICE_ADDRESS" envDefault:"key-service.k8sdeploy:8001"`
}

type Services struct {
	OrchestratorService
	KeyService
}

type Local struct {
	KeepLocal   bool `env:"LOCAL_ONLY" envDefault:"false" json:"keep_local,omitempty"`
	Development bool `env:"DEVELOPMENT" envDefault:"false" json:"development,omitempty"`
	HTTPPort    int  `env:"HTTP_PORT" envDefault:"3000" json:"port,omitempty"`
	Services    `json:"services"`
}

func buildServiceKeys(cfg *Config) error {
	vaultSecrets, err := cfg.getVaultSecrets("kv/data/k8sdeploy/api-keys")
	if err != nil {
		return err
	}
	secrets, err := ParseKVSecrets(vaultSecrets)
	if err != nil {
		return err
	}

	for _, secret := range secrets {
		if secret.Key == "hooks" {
			cfg.Local.KeyService.Key = secret.Value
			cfg.Local.OrchestratorService.Key = secret.Value
		}
	}

	return nil
}

func BuildLocal(cfg *Config) error {
	local := &Local{}
	if err := env.Parse(local); err != nil {
		return err
	}
	cfg.Local = *local

	if err := buildServiceKeys(cfg); err != nil {
		return err
	}

	return nil
}
