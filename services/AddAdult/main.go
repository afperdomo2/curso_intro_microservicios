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

	r.POST("/Add/Adults", AddAdult)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// AddAdult godoc
// @Summary Add a new adult
// @Description Add a new adult
// @Tags Adults
// @Accept  json
// @Produce  json
// @Param adult body models.Adult true "Adult Object"
// @Success 200 {object} models.Response
// @Router /Add/Adults [post]
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
