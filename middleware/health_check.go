package middleware

import (
	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
)

func SetupHealthcheck(r *gin.Engine) {
	r.Use(healthcheck.Default())
}
