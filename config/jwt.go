package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("jkhdkfjghdlkfghhjdbfggsedrjklbg")

type JWTClain struct {
	Username string
	jwt.RegisteredClaims
}
