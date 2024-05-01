package config

import (
	"github.com/doxanocap/pkg/env"
	"github.com/doxanocap/pkg/sandbox/lg"
)

const (
	EnvProduction  = "prod"
	EnvDevelopment = "dev"
	EnvStage       = "stage"
)

var (
	DefaultEnvPath = ".env"
)

type IConfig interface {
	Env() string
}

func InitConfig[T IConfig]() *T {
	var config T

	err := env.LoadFile(DefaultEnvPath)
	if err != nil {
		lg.Errorf("config: %s", err)
	}

	err = env.Unmarshal(&config)
	if err != nil {
		lg.Fatalf("config: %s", err)
	}

	return &config
}
