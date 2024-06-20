package logger

import (
	"github.com/doxanocap/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	log *zap.Logger
}

func InitLogger[C config.IConfig](cfg *C) *zap.Logger {
	writer := zapcore.Lock(os.Stdout)
	encoder := getEncoder(cfg)
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	return zap.New(core)
}

func getEncoder[C config.IConfig](cfg *C) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{}
	if (*cfg).Env() == config.EnvProduction {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"

	return zapcore.NewJSONEncoder(encoderConfig)
}
