package logger

import (
	"context"
	"fmt"
	"github.com/doxanocap/pkg/config"
	"log/slog"
	"os"
)

var (
	defaultHandler slog.Handler
)

const (
	keyPayload = "payload"
	keyModule  = "module"
)

type SlogLogger struct {
	log    *slog.Logger
	attrs  []slog.Attr
	module string
}

// InitSlogLogger
//
//	{
//		"time":"2024-05-20T14:11:29.7812801+05:00",
//		"level":"ERROR",
//		"msg":"unable to unmarshal",
//		"module": "[REPOSITORY][USER_CACHE]",
//		"payload": {
//	  	"incoming_request": {
//				"method": "POST",
//				"path": "/v1/auth/sign-in",
//				"latency": "120ms",
//				"status": 404,
//				"body": "{"\message"\: "\not found"\}"
//			},
//			"msg_id": "312de-423f3j-ew4043-wf43",
//			"send_to": 325,
//		}
//	}
func InitSlogLogger(env string) *SlogLogger {
	switch env {
	case config.EnvDevelopment:
		defaultHandler = NewColorfulHandler(os.Stdout, nil)
	case config.EnvStage:
		defaultHandler = slog.NewTextHandler(os.Stdout, nil)
	case config.EnvProduction:
		defaultHandler = slog.NewJSONHandler(os.Stdout, nil)
	}

	return &SlogLogger{
		log: slog.New(defaultHandler),
	}
}

func (l *SlogLogger) WithGroup(name string) *SlogLogger {
	return &SlogLogger{
		log:    l.log.WithGroup(name),
		module: l.module,
	}
}

func (l *SlogLogger) WithModule(module string) *SlogLogger {
	return &SlogLogger{
		log:    l.log,
		module: fmt.Sprintf("%s[%s]", l.module, module),
	}
}

func (l *SlogLogger) WithAttrs(attrs ...slog.Attr) *SlogLogger {
	return &SlogLogger{
		attrs:  attrs,
		log:    l.log,
		module: l.module,
	}
}

func (l *SlogLogger) Log(level slog.Level, msg string, attrs ...slog.Attr) {
	if len(l.attrs) != 0 {
		attrs = append(l.attrs, attrs...)
	}
	logAttrs := []slog.Attr{
		{
			Key:   keyPayload,
			Value: slog.GroupValue(attrs...),
		},
	}
	if l.module != "" {
		logAttrs = append(logAttrs, slog.String(keyModule, l.module))
	}

	l.log.LogAttrs(context.Background(), level, msg, logAttrs...)
}

func (l *SlogLogger) Info(msg string, args ...slog.Attr) {
	l.Log(slog.LevelInfo, msg, args...)
}

func (l *SlogLogger) Debug(msg string, args ...slog.Attr) {
	l.Log(slog.LevelDebug, msg, args...)
}

func (l *SlogLogger) Warn(msg string, args ...slog.Attr) {
	l.Log(slog.LevelWarn, msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...slog.Attr) {
	l.Log(slog.LevelError, msg, args...)
}
