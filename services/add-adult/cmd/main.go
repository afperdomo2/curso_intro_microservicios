package main

import (
	"afperdomo2/go/microservicios/pkg/database"
	"afperdomo2/go/microservicios/services/add-adult/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.POST("/Add/Adults", handlers.AddAdult)

	r.Run(":8080")
}
