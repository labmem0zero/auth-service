package usecases

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Crypto interface {
	GenerateUserToken(username string) (bearer string, err error)
}

type CryptoHS256 struct {
	privateKey []byte
}

func (c *CryptoHS256) GenerateUserToken(username string) (bearer string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(15 * time.Minute).Unix(),
		})
	return token.SignedString(c.privateKey)
}
