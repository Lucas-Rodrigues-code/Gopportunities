package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize()  {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start the server on port 8080
	r.Run(":8080") // Listen and serve on	
}