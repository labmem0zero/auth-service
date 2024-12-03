package dbrepo

import (
	"github.com/jmoiron/sqlx"
	"github.com/labmem0zero/go-logger"

	"auth/models"
)

type UserDBRepo interface {
	UserSet(reqID string, u models.UserSet) (res int64, err error)
}

func NewUserDBRepo(l *logger.Logger, db *sqlx.DB) UserDBRepo {
	return userDBRepo{
		Logger: l,
		udb:    db,
	}
}

type userDBRepo struct {
	*logger.Logger
	udb *sqlx.DB
}

func (d userDBRepo) UserSet(reqID string, u models.UserSet) (res int64, err error) {
	q := `SELECT * FROM auth.user_set($1,$2,$3)`
	d.Info(reqID, q, u)
	if err = d.udb.Get(&res, q, u.UserID, u.Username, u.Password); err != nil {
		d.Error(reqID, err)
	}

	return
}
