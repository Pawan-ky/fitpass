package main

import (
	"fmt"
	db "fitpass/database"
	"github.com/gin-gonic/gin"
	"fitpass/routes"
)

func main() {
    fmt.Println("hello")

	db.Init()
	db.Migrate()

	router := gin.Default()
	api := router.Group("/api")
	routes.AllRoutes(api)

	router.Run(":8080")
}

