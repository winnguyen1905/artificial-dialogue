package main

import (
	"artificial-dialogue/cmd/cli"
	"artificial-dialogue/internal/router"
)

func main() {
	
	r := router.IndexRouter()

	r.Run(":8080")
}