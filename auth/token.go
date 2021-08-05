package auth

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(user_id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func TokenValid(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
}

func ExtractToken(c *gin.Context) (string, error) {
	const BEARER_SCHEMA = "Bearer"
	if authHeader := c.GetHeader("Authorization"); authHeader != "" {
		return authHeader[len(BEARER_SCHEMA)+1:], nil
	}
	return "", fmt.Errorf("unauthorized")
}
