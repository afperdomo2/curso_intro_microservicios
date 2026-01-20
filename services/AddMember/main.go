package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// No Database connection for now
	r := gin.Default()

	r.POST("/Add/Member", AddMember)

	r.Run(":8080")
}

func AddMember(c *gin.Context) {
	// Dummy implementation not connected to DB
	c.JSON(http.StatusOK, gin.H{
		"message": "Member received successfully (Not saved to DB)",
	})
}
