package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"fitpass/routes"
)

func main() {
    fmt.Println("hello")

	router := gin.Default()
	api := router.Group("/api")
	routes.AllRoutes(api)

	router.Run(":8080")
}

