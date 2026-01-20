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

	r.GET("/Adults/:id", GetAdultById)

	r.Run(":8080")
}

func GetAdultById(c *gin.Context) {
	id := c.Param("id")
	var adult models.Adult
	if err := database.DB.First(&adult, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Adult not found"})
		return
	}
	c.JSON(http.StatusOK, adult)
}
