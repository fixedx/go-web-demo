package main

import (
	"go-web-demo/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
