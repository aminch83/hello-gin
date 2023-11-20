package server

import (
	"hello-gin/src/api/health"
	"hello-gin/src/api/products"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func NewRouter() *gin.Engine {
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health.InitRoute(router)
	products.InitRoute(router)

	return router
}
