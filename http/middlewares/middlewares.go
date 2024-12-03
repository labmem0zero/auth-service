package middlewares

import (
	"crypto/rsa"

	"github.com/labmem0zero/go-logger"
)

type Middlewares struct {
	l   *logger.Logger
	pub *rsa.PublicKey
}

func New(l *logger.Logger, pub *rsa.PublicKey) Middlewares {
	return Middlewares{
		l:   l,
		pub: pub,
	}
}
