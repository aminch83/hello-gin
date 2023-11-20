// create the same route as products for users
package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// create a Struct for user with id, name, and email
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// add a new route to return a list of users
func getUsers(c *gin.Context) {
	// create a slice of users
	users := []User{
		{ID: 1, Name: "John Doe", Email: "johndoe@live.com"},
		{ID: 2, Name: "Jane Doe", Email: ""},
		{ID: 3, Name: "John Smith", Email: ""},
	}

	// return the list of users
	c.JSON(http.StatusOK, users)
}

// add a new route to return a single user
func getUser(c *gin.Context) {
	// create a slice of users
	user := User{ID: 1, Name: "John Doe", Email: "johndoe@live.com"}

	// return the list of users
	c.JSON(http.StatusOK, user)
}

// create the router with the user handlers
func HandleUsers(router *gin.Engine) {
	router.GET("/users", getUsers)
	router.GET("/user/:id", getUser)
}
