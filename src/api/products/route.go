package products

import (
	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.Engine) {
	controller := ProductController{}
	router.GET("/products", controller.getProducts)
	router.GET("/products/", controller.getProducts)
	router.GET("/products/:id", controller.getProduct)
}
