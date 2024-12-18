package usecases

import (
	"github.com/labmem0zero/go-logger"

	"auth/models"
	"auth/provider"
	"auth/services/usecases/dbrepo"
	"auth/util"
)

type Usecases struct {
	l *logger.Logger
	*util.Recovery
	udbr dbrepo.UserDBRepo
}

func New(p provider.Provider, l *logger.Logger, r *util.Recovery) Usecases {
	return Usecases{
		udbr:     dbrepo.NewUserDBRepo(l, p.GetUsersDB()),
		l:        l,
		Recovery: r,
	}
}

func (u Usecases) Start(reqID string) {
	u.l.Info(reqID, "Usecases has started")
}

func (u Usecases) Stop(reqID string) {
	u.l.Info(reqID, "Usecases has stopped")
}

func (u Usecases) UserCreate(reqID string, req models.UserSet) (res models.UserView, err error) {
	var id int64
	id, err = u.udbr.UserSet(reqID, req)
	res = models.UserView{UserID: id, Username: req.Username}
	return
}
