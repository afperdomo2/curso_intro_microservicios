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

	r.POST("/Add/Adults", AddAdult)

	r.Run(":8080")
}

func AddAdult(c *gin.Context) {
	var adult models.Adult
	if err := c.ShouldBindJSON(&adult); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&adult).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Message: "Adult added successfully"})
}
