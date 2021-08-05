package middleware

import (
	"time"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "OPTION", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
}

func SetupHealthcheck(r *gin.Engine) {
	r.Use(healthcheck.Default())
}
