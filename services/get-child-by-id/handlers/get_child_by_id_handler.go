package handlers

import (
	"net/http"

	"afperdomo2/go/microservicios/pkg/database"
	"afperdomo2/go/microservicios/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetChildById(c *gin.Context) {
	id := c.Param("id")
	var child models.Child
	if err := database.DB.First(&child, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Child not found"})
		return
	}
	c.JSON(http.StatusOK, child)
}
