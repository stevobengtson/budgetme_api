package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtCustomClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func CreateToken(user_id uint) (string, error) {
	var claims JwtCustomClaims
	claims.UserId = user_id
	claims.ExpiresAt = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("BUDGETME_USER_API_SECRET")))
}

func ExtractToken(c *gin.Context) (string, error) {
	const BEARER_SCHEMA = "Bearer"
	if authHeader := c.GetHeader("Authorization"); authHeader != "" {
		return authHeader[len(BEARER_SCHEMA)+1:], nil
	}
	return "", fmt.Errorf("unauthorized")
}

func IsTokenValid(c *gin.Context) (uint, error) {
	var tokenString string
	var err error

	if tokenString, err = ExtractToken(c); err != nil {
		return 0, errors.New("invalid token")
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("BUDGETME_USER_API_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return 0, errors.New("unable to extract token")
}
