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

	r.GET("/Children/:id", GetChildById)

	r.Run(":8080")
}

func GetChildById(c *gin.Context) {
	id := c.Param("id")
	var child models.Child
	if err := database.DB.First(&child, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Child not found"})
		return
	}
	c.JSON(http.StatusOK, child)
}
