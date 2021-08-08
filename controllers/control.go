package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevobengtson/budgetme/seed"
)

func SeedData(c *gin.Context) {
	if err := seed.Load(); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
