package dbrepo

import (
	"github.com/jmoiron/sqlx"
	"github.com/labmem0zero/go-logger"

	"auth/models"
)

type UserDBRepo interface {
	UserCreate(reqID string, user models.UserCreate) (res int64, err error)
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

func (d userDBRepo) UserCreate(reqID string, user models.UserCreate) (res int64, err error) {
	err = d.udb.Get(&res, `
	INSERT INTO vpn_auth_service.users(username, pw_hash)
	VALUES ($1, $2)
	ON CONFLICT (username) DO NOTHING
	RETURNING user_id`,
		user.Username,
		user.PWHash)
	return
}
