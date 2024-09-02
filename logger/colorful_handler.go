package logger

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"

	"github.com/fatih/color"
)

type ColorfulHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type ColorfulHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *ColorfulHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		if a.Key == keyModule {

		}
		fields[a.Key] = a.Value.Any()
		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[15:04:05.000 -0700]")

	h.l.Println(timeStr, level, r.Message, color.WhiteString(string(b)))
	return nil
}

func NewColorfulHandler(out io.Writer, opts *slog.HandlerOptions) *ColorfulHandler {
	return &ColorfulHandler{
		Handler: slog.NewJSONHandler(out, opts),
		l:       log.New(out, "", 0),
	}
}
