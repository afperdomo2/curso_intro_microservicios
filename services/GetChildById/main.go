package main

import (
	"net/http"

	"afperdomo2/go/microservicios/pkg/database"
	"afperdomo2/go/microservicios/pkg/models"

	_ "afperdomo2/go/microservicios/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ApiGlobal
// @version 1.0
// @description API for managing Children
// @host localhost:8080
// @BasePath /

func main() {
	database.InitDB()
	r := gin.Default()

	r.GET("/Children/:id", GetChildById)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// GetChildById godoc
// @Summary Get child by ID
// @Description Get a specific child by ID
// @Tags Children
// @Produce  json
// @Param id path string true "Child ID"
// @Success 200 {object} models.Child
// @Router /Children/{id} [get]
func GetChildById(c *gin.Context) {
	id := c.Param("id")
	var child models.Child
	if err := database.DB.First(&child, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Child not found"})
		return
	}
	c.JSON(http.StatusOK, child)
}
