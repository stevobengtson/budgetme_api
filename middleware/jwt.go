package middleware

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stevobengtson/budgetme_api/auth"
	"github.com/stevobengtson/budgetme_api/models"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string
		var token *jwt.Token
		var err error

		if tokenString, err = auth.ExtractToken(c); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if token, err = auth.TokenValid(tokenString); err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if expiry, ok := claims["exp"].(int64); !ok || expiry >= time.Now().Unix() {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			}

			if userId, ok := claims["user_id"].(uint); ok {
				var user models.User
				if err = models.GetUserByID(&user, userId); err == nil {
					c.AbortWithStatus(http.StatusUnauthorized)
				}

				c.Set("currentUser", user)
				return
			}
		}

		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
