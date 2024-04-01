package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type HSSignMethod string

const HS256 HSSignMethod = "HS256"

type SignCliams struct {
	jwt.RegisteredClaims
	UserId int64
}

func NewCliams(userId int64) SignCliams {
	return SignCliams{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
}

type HS struct {
	Key        string
	SignMethod HSSignMethod
}

func NewHS() *HS {
	return &HS{
		Key:        "jwt-gf-naive",
		SignMethod: HS256,
	}
}
func (hs *HS) getSignMethod() *jwt.SigningMethodHMAC {
	switch hs.SignMethod {
	case HS256:
		return jwt.SigningMethodHS256
	}
	return jwt.SigningMethodHS256
}
func (hs *HS) Encode(claims SignCliams) (string, error) {
	withClaims := jwt.NewWithClaims(hs.getSignMethod(), claims)
	signingString, err := withClaims.SignedString([]byte(hs.Key))
	return signingString, err
}
func (hs *HS) Decode(sign string, claims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(sign, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(hs.Key), nil
	})
	return err
}
