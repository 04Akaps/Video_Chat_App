package reposiroty

import (
	"encoding/base64"
	"github.com/o1egl/paseto"
	"math/rand"
	"time"
)

type PasetoMaker struct {
	Paseto *paseto.V2
	Key    []byte
}

type payload struct {
	UserName     string
	CurrentTime  time.Time
	RandomString string
}

func NewPasetoMaker(key string) *PasetoMaker {
	return &PasetoMaker{
		Paseto: paseto.NewV2(),
		Key:    []byte(key),
	}
}

func (maker *PasetoMaker) CreateToken(userName string) (string, error) {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)

	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	newPayload := &payload{
		CurrentTime:  time.Now(),
		RandomString: randomString,
		UserName:     userName,
	}

	return maker.Paseto.Encrypt(maker.Key, newPayload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (string, error) {
	p := &payload{}
	err := maker.Paseto.Decrypt(token, maker.Key, p, nil)

	if err != nil {
		return "", err
	}
	return p.UserName, nil
}
