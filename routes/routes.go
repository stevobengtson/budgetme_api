package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stevobengtson/budgetme_api/controllers"
	"github.com/stevobengtson/budgetme_api/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	loginGrp := r.Group("/auth")
	{
		loginGrp.POST("/", controllers.Login)
	}

	grp1 := r.Group("/user")
	{
		grp1.GET("/", middleware.AuthorizeJWT(), controllers.GetUsers)
		grp1.POST("/", controllers.CreateUser)
		grp1.GET("/:id", middleware.AuthorizeJWT(), controllers.GetUserByID)
		grp1.PUT("/:id", middleware.AuthorizeJWT(), controllers.UpdateUser)
		grp1.DELETE("/:id", middleware.AuthorizeJWT(), controllers.DeleteUser)
	}
	return r
}
