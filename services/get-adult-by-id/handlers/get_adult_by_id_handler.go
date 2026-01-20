package handlers

import (
	"net/http"

	"afperdomo2/go/microservicios/pkg/database"
	"afperdomo2/go/microservicios/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetAdultById(c *gin.Context) {
	id := c.Param("id")
	var adult models.Adult
	if err := database.DB.First(&adult, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Adult not found"})
		return
	}
	c.JSON(http.StatusOK, adult)
}
