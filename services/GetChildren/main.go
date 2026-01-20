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

	r.GET("/Children", GetChildren)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// GetChildren godoc
// @Summary Get all children
// @Description Get a list of children
// @Tags Children
// @Produce  json
// @Success 200 {array} models.Child
// @Router /Children [get]
func GetChildren(c *gin.Context) {
	var children []models.Child
	if err := database.DB.Find(&children).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, children)
}
