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
// @description API for managing Adults
// @host localhost:8080
// @BasePath /

func main() {
	database.InitDB()
	r := gin.Default()

	r.GET("/Adults/:id", GetAdultById)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// GetAdultById godoc
// @Summary Get adult by ID
// @Description Get a specific adult by ID
// @Tags Adults
// @Produce  json
// @Param id path string true "Adult ID"
// @Success 200 {object} models.Adult
// @Router /Adults/{id} [get]
func GetAdultById(c *gin.Context) {
	id := c.Param("id")
	var adult models.Adult
	if err := database.DB.First(&adult, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Adult not found"})
		return
	}
	c.JSON(http.StatusOK, adult)
}
