package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevobengtson/budgetme/auth"
	"github.com/stevobengtson/budgetme/models"
	"github.com/stevobengtson/budgetme/repository"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userId uint
		var err error
		var user models.User

		if userId, err = auth.IsTokenValid(c); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if err = repository.GetUserByID(&user, userId); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("currentUser", user)
	}
}
