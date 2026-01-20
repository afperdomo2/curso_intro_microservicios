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

	r.GET("/Adults", GetAdults)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// GetAdults godoc
// @Summary Get all adults
// @Description Get a list of adults
// @Tags Adults
// @Produce  json
// @Success 200 {array} models.Adult
// @Router /Adults [get]
func GetAdults(c *gin.Context) {
	var adults []models.Adult
	if err := database.DB.Find(&adults).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adults)
}
