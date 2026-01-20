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

	r.POST("/Add/Children", AddChild)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// AddChild godoc
// @Summary Add a new child
// @Description Add a new child
// @Tags Children
// @Accept  json
// @Produce  json
// @Param child body models.Child true "Child Object"
// @Success 200 {object} models.Response
// @Router /Add/Children [post]
func AddChild(c *gin.Context) {
	var child models.Child
	if err := c.ShouldBindJSON(&child); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&child).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Message: "Child added successfully"})
}
