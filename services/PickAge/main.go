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

	r.GET("/PickAge", PickAge)

	r.Run(":8080")
}

func PickAge(c *gin.Context) {
	// Reusing Adult model logic as a placeholder for PickAge functionality
	// assuming it returns a list of adults for now as requested "like the others"
	var adults []models.Adult
	if err := database.DB.Find(&adults).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adults)
}
