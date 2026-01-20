package main

import (
	"afperdomo2/go/microservicios/pkg/database"
	"afperdomo2/go/microservicios/services/get-children/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.GET("/Children", handlers.GetChildren)

	r.Run(":8080")
}
