package router

import (
	"fmt"

	"github.com/Lucas-Rodrigues-code/Gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	fmt.Println("Initializing router...")
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/opening", handler.ListOpeningsHandler)
	}
}
