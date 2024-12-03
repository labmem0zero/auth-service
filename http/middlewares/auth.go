package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	"auth/errs"
	http2 "auth/http"
)

func (mw *Middlewares) MWAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ah := r.Header.Get("Authorization")
		split := strings.Split(ah, " ")
		if len(split) != 2 || split[0] != "Bearer" {
			http2.CheckErrWriteResp(w, 401, nil, errs.New(errors.New("No auth token"), "Authorization required", 401))
			return
		}
		//claims, err := jwt.Parse(split[1], mw.keyfunc)
		h.ServeHTTP(w, r)
	})
}

func (mw *Middlewares) keyfunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, errors.New("Parse token error")
	}

	return mw.pub, nil
}
