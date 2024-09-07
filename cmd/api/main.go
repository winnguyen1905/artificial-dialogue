package main

import (
	// "artificial-dialogue/internal/controller"
	// "fmt"
	// "artificial-dialogue/client"
	// "artificial-dialogue/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// bot := client.Init()

 	// handler.Init(bot)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}