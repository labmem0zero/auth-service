package handlers

import (
	"github.com/labmem0zero/go-logger"

	"auth/services/usecases"
)

type Handlers struct {
	l *logger.Logger
	u usecases.Usecases
}

func New(l *logger.Logger, u usecases.Usecases) Handlers {
	return Handlers{
		l: l,
		u: u,
	}
}
