package main

import (
	"afperdomo2/go/microservicios/pkg/database"
	"afperdomo2/go/microservicios/services/get-adults/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.GET("/Adults", handlers.GetAdults)

	r.Run(":8080")
}
