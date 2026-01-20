package main

import (
	"afperdomo2/go/microservicios/services/add-member/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// No Database connection for now
	r := gin.Default()

	r.POST("/Add/Member", handlers.AddMember)

	r.Run(":8080")
}
