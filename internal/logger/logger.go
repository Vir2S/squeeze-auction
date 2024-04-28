package logger

import (
	log "github.com/sirupsen/logrus"
	"squeeze-auction/internal/config"
)

type Logger struct {
	*log.Logger
}

func New(cfg *config.Config) *Logger {
	l := log.New()

	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Warnf("Invalid log level '%s'. Using 'info' level as default.", cfg.LogLevel)
		level = log.InfoLevel
	}
	l.SetLevel(level)

	return &Logger{l}
}
