package main

import (
	"afperdomo2/go/microservicios/pkg/database"
	"afperdomo2/go/microservicios/services/pick-age/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.GET("/PickAge", handlers.PickAge)

	r.Run(":8080")
}
