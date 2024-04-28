package routes

import (
	"squeeze-auction/internal/logger"
)

type Routes struct {
	Log *logger.Logger
}

func New(log *logger.Logger) *Routes {
	return &Routes{
		Log: log,
	}
}
