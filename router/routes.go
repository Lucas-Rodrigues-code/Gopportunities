package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsersHandler(c *gin.Context) {
	// Example response for the /users endpoint
	users := []string{"Alice", "Bob", "Charlie"}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
func initializeRoutes(r *gin.Engine) {
	fmt.Println("Initializing router...")
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", getUsersHandler)
	}
}
