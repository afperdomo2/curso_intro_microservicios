package main

import (
	"net/http"

	_ "afperdomo2/go/microservicios/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ApiGlobal
// @version 1.0
// @description API for managing Adults and Children
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	r.GET("/Adults", GetAdults)
	r.GET("/Children", GetChildren)
	r.GET("/Adults/:id", GetAdultById)
	r.GET("/Children/:id", GetChildById)
	r.POST("/Add/Adults", AddAdult)
	r.POST("/Add/Children", AddChild)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// Response represents a generic response
type Response struct {
	Message string `json:"message"`
}

// GetAdults godoc
// @Summary Get all adults
// @Description Get a list of adults
// @Tags Adults
// @Produce  json
// @Success 200 {object} Response
// @Router /Adults [get]
func GetAdults(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "Adults list retrieved successfully"})
}

// GetChildren godoc
// @Summary Get all children
// @Description Get a list of children
// @Tags Children
// @Produce  json
// @Success 200 {object} Response
// @Router /Children [get]
func GetChildren(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "Children list retrieved successfully"})
}

// GetAdultById godoc
// @Summary Get adult by ID
// @Description Get a specific adult by ID
// @Tags Adults
// @Produce  json
// @Param id path string true "Adult ID"
// @Success 200 {object} Response
// @Router /Adults/{id} [get]
func GetAdultById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, Response{Message: "Adult retrieved successfully with ID: " + id})
}

// GetChildById godoc
// @Summary Get child by ID
// @Description Get a specific child by ID
// @Tags Children
// @Produce  json
// @Param id path string true "Child ID"
// @Success 200 {object} Response
// @Router /Children/{id} [get]
func GetChildById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, Response{Message: "Child retrieved successfully with ID: " + id})
}

// AddAdult godoc
// @Summary Add a new adult
// @Description Add a new adult
// @Tags Adults
// @Produce  json
// @Success 200 {object} Response
// @Router /Add/Adults [post]
func AddAdult(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "Adult added successfully"})
}

// AddChild godoc
// @Summary Add a new child
// @Description Add a new child
// @Tags Children
// @Produce  json
// @Success 200 {object} Response
// @Router /Add/Children [post]
func AddChild(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "Child added successfully"})
}
