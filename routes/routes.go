package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stevobengtson/budgetme/controllers"
	"github.com/stevobengtson/budgetme/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	loginGrp := r.Group("/auth")
	{
		loginGrp.POST("/", controllers.Login)
		loginGrp.PUT("/refresh", middleware.AuthorizeJWT(), controllers.Refresh)
	}

	userGroup := r.Group("/user")
	{
		userGroup.GET("/", middleware.AuthorizeJWT(), controllers.GetUsers)
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/:id", middleware.AuthorizeJWT(), controllers.GetUserByID)
		userGroup.PUT("/:id", middleware.AuthorizeJWT(), controllers.UpdateUser)
		userGroup.DELETE("/:id", middleware.AuthorizeJWT(), controllers.DeleteUser)
	}

	accountGrp := r.Group("/account")
	{
		accountGrp.GET("/", middleware.AuthorizeJWT(), controllers.GetAccounts)
		accountGrp.POST("/", middleware.AuthorizeJWT(), controllers.CreateAccount)
	}

	devGroup := r.Group("/dev")
	{
		devGroup.GET("/seed", controllers.SeedData)
	}

	return r
}
