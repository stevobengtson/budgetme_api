package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/stevobengtson/user_service/models"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	pagination := models.GeneratePaginationFromRequest(c)
	err := models.GetAllUsersPaged(&users, &pagination)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
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
		return
	}

	if currentUser := c.Value("currentUser").(models.User); currentUser.ID != user.ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unable to update other users"})
		return
	}

	c.BindJSON(&user)
	err = models.UpdateUser(&user, fmt.Sprint(user.ID))
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
		return
	}

	if currentUser := c.Value("currentUser").(models.User); currentUser.ID != user.ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unable to delete other users"})
		return
	}

	err = models.DeleteUser(&user, fmt.Sprint(user.ID))
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
