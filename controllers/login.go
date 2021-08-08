package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevobengtson/budgetme/auth"
	"github.com/stevobengtson/budgetme/models"
	"github.com/stevobengtson/budgetme/repository"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	err := user.Validate("login")
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	token, err := SignIn(user.Email, user.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Refresh(c *gin.Context) {
	currentUser := c.Value("currentUser").(models.User)

	token, err := auth.CreateToken(currentUser.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func SignIn(email, password string) (string, error) {
	var user models.User

	err := repository.GetUserByEmail(&user, email)
	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(user.ID)
}
