package health

import (
	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.Engine) {
	controller := HealthController{}
	router.GET("/health", controller.healthCheck)
	router.GET("/health/", controller.healthCheck)
}
