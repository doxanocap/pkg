package logger

import (
	"context"
	"fmt"
	"github.com/doxanocap/pkg/config"
	"log/slog"
	"os"
)

const (
	keyPayload = "payload"
	keyModule  = "module"
)

type SlogLogger struct {
	log    *slog.Logger
	module string
}

//{
//	"time":"2024-05-20T14:11:29.7812801+05:00",
//	"level":"ERROR",
//	"msg":"unable to unmarshal",
//	"module": "[REPOSITORY][USER_CACHE]",
//	"payload": {
//		"msg_id": "312de-423f3j-ew4043-wf43",
//		"send_to": 325,
//	}
//}

func InitSlogLogger(env string) *SlogLogger {
	var handler slog.Handler
	if env == config.EnvDevelopment {
		handler = slog.NewTextHandler(os.Stdout, nil)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	}

	return &SlogLogger{
		log: slog.New(handler),
	}
}

func (l *SlogLogger) WithModule(module string) *SlogLogger {
	return &SlogLogger{
		log:    l.log,
		module: fmt.Sprintf("%s[%s]", l.module, module),
	}
}

func (l *SlogLogger) Info(msg string, args ...slog.Attr) {
	attrs := []slog.Attr{{Key: keyPayload, Value: slog.GroupValue(args...)}}
	if l.module != "" {
		attrs = append(attrs, slog.String(keyModule, l.module))
	}

	l.log.LogAttrs(context.Background(),
		slog.LevelInfo,
		msg,
		attrs...)
}
