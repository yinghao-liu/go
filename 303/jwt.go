package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Minute)
	issuer := username
	claims := Claims{
		ID:       10001,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}

func ParseToken(token string, claims jwt.Claims) (*jwt.Token, error) {

	tok, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	return tok, err
}
