package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/stevobengtson/budgetme_api/models"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	user, err := getUserFromParams(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	user, err := getUserFromParams(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if currentUser := c.Value("currentUser").(models.User); currentUser.Id != user.Id {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unable to update other users"})
	}

	c.BindJSON(&user)
	err = models.UpdateUser(&user, fmt.Sprint(user.Id))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	user, err := getUserFromParams(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if currentUser := c.Value("currentUser").(models.User); currentUser.Id != user.Id {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unable to delete other users"})
	}

	err = models.DeleteUser(&user, fmt.Sprint(user.Id))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "user is deleted"})
	}
}

func getUserFromParams(c *gin.Context) (models.User, error) {
	var user models.User
	var userId uint64
	var err error

	userId, err = strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	if err != nil {
		return user, err
	}

	err = models.GetUserByID(&user, uint(userId))
	if err != nil {
		return user, err
	}

	return user, nil
}
