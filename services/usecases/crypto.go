package usecases

type Crypto interface {
	Generate() (public, secret []byte, err error)
}
