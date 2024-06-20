package logger

import (
	"testing"
)

type T struct {
	env string
}

func (t T) Env() string {
	return t.env
}

func TestLogger(t *testing.T) {
	//log := InitLogger(config.EnvProduction)
	//
	//repoLogger := log.WithModule("REPOSITORY")
	//
	//userRepoLogger := repoLogger.WithModule("USER")
	//
	//userRepoLogger.Info("updated counter",
	//	slog.String("id", "e32me-r23mf9-gg34r3-dssd3e"),
	//	slog.Int("counter", 3123),
	//)
}
