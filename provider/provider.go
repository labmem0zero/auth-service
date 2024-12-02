package provider

import (
	"github.com/jmoiron/sqlx"
	"github.com/labmem0zero/go-logger"

	"auth/config"
	"auth/provider/db"
)

type Provider struct {
	l       *logger.Logger
	usersDB *sqlx.DB
}

func New(conf config.Config, l *logger.Logger) (p Provider, err error) {
	p = Provider{l: l}
	if p.usersDB, err = db.Connect(conf.UsersDB); err != nil {
		return
	}
	return
}

func (p Provider) GetUsersDB() *sqlx.DB {
	return p.usersDB
}
