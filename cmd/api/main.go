package main

import (
	"artificial-dialogue/internal/router"
)

func main() {
	
	r := router.IndexRouter()

	r.Run(":8080")
}