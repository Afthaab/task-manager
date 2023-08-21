package domain

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	Userid uint
	Source string
	Role   string
	jwt.StandardClaims
}
