package pkg

import (
	"fmt"
	"github.com/doxanocap/pkg/config"
	"github.com/doxanocap/pkg/logger"
	"github.com/doxanocap/pkg/router"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

type TConfig struct {
	ENV        string `env:"APP_ENV"`
	ServerPORT string `env:"SERVER_PORT"`
}

func (c TConfig) Env() string {
	return c.ENV
}

func Test_PkgDeps(t *testing.T) {
	cfg := config.InitConfig[TConfig]()

	fmt.Println(cfg.ServerPORT, cfg.ENV)

	log := logger.InitLogger[TConfig](cfg).Named("[TEST]")

	//assert.Equal(t, cfg.ENV, "prod")
	assert.Equal(t, cfg.ServerPORT, "5000")
	assert.NotEmpty(t, log)

	r := router.InitGinRouter(cfg.ENV)
	assert.NotEmpty(t, r)

	log = log.With(zap.Any("payload",
		[]zap.Field{
			zap.String("key", "hello"),
			zap.String("value", "world!"),
		}))

	log.Info("ok!")
}
