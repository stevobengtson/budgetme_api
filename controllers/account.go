package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevobengtson/budgetme/models"
	"github.com/stevobengtson/budgetme/repository"
)

func CreateAccount(c *gin.Context) {
	var account models.Account
	user := models.GetCurrentUser(c)
	if user == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.BindJSON(&account)
	account.UserId = user.ID

	err := repository.CreateAccount(&account)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": account})
	}
}

func GetAccounts(c *gin.Context) {
	var accounts []models.Account
	user := models.GetCurrentUser(c)
	if user == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	err := repository.GetAllAccounts(user.ID, &accounts)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": accounts})
	}
}
