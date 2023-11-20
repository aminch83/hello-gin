package products

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

// create a Struct for product with id, name, and price
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// add a new route to return a list of products
func (controller *ProductController) getProducts(c *gin.Context) {
	// create a slice of products
	products := []Product{
		{ID: 1, Name: "laptop", Price: 1000},
		{ID: 2, Name: "mouse", Price: 50},
		{ID: 3, Name: "keyboard", Price: 75},
	}

	// return the list of products
	c.JSON(http.StatusOK, products)
}

// add a new route to return a single product
func (controller *ProductController) getProduct(c *gin.Context) {
	log.Println("Product Id: ", c.Param("id"))
	// create a slice of products
	product := Product{ID: 1, Name: "laptop", Price: 1000}

	// return the list of products
	c.JSON(http.StatusOK, product)
}
