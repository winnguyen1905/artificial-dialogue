package router

import (
	"artificial-dialogue/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexRouter() *gin.Engine {
	r := gin.Default()
	
	api := r.Group("/api/v1") 
	{
		api.GET("/", controller.NewUserController().GetUserById)
	}

	return r
}

func Pong (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}