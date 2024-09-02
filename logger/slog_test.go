package logger

import (
	"github.com/doxanocap/pkg/config"
	"log/slog"
	"testing"
)

func TestLogger(t *testing.T) {
	log := InitSlogLogger(config.EnvDevelopment)

	repoLogger := log.WithModule("REPOSITORY")

	userRepoLogger := repoLogger.WithModule("USER")

	userRepoLogger.Info("updated counter",
		slog.String("id", "e32me-r23mf9-gg34r3-dssd3e"),
		slog.Int("counter", 3123),
	)

	userRepoLogger.Error("timeout",
		slog.String("id", "e32me-r23mf9-gg34r3-dssd3e"),
		slog.Int("counter", 3123),
	)
}
