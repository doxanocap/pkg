package config

import (
	"github.com/doxanocap/pkg/env"
	"log"
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
		log.Printf("error: config: %s", err)
	}

	err = env.Unmarshal(&config)
	if err != nil {
		log.Fatalf("config: %s", err)
	}

	return &config
}
