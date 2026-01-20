package main

import (
	"net/http"

	"afperdomo2/go/microservicios/pkg/database"
	"afperdomo2/go/microservicios/pkg/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.GET("/Children", GetChildren)

	r.Run(":8080")
}

func GetChildren(c *gin.Context) {
	var children []models.Child
	if err := database.DB.Find(&children).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, children)
}
