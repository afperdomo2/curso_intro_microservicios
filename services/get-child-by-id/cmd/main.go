package main

import (
	"afperdomo2/go/microservicios/pkg/database"
	"afperdomo2/go/microservicios/services/get-child-by-id/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.GET("/Children/:id", handlers.GetChildById)

	r.Run(":8080")
}
