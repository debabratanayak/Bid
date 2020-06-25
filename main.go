package main

import (
	"test1/routes"

	"github.com/gin-gonic/gin"
)

var err error

func main() {
	router := gin.Default()
	routes.IntializeRoutes(router)
	router.Run(":8080")
}
