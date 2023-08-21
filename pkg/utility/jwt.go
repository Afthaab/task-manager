package utility

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/afthaab/task-manager/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetTheBearerToken(c *gin.Context) (string, error) {
	authorization := c.Request.Header.Get("authorization")

	if authorization == "" {

		return "", errors.New("Nil value in the authorization")
	}
	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		return "", errors.New("Bearer string is broken")
	}
	return token[1], nil
}

func GetTokenFromString(signedToken string, claims *domain.JWTClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KET")), nil
	})
}
