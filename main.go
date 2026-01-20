package main

import (
	"fmt"
	"log"
	"net/http"

	_ "afperdomo2/go/microservicios/docs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title ApiGlobal
// @version 1.0
// @description API for managing Adults and Children
// @host localhost:8080
// @BasePath /

// Global DB variable
var DB *gorm.DB

// Models
type Adult struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id" swaggertype:"string" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name      string    `json:"name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	BirthYear int       `json:"birth_year" example:"1990"`
	ImageURL  string    `json:"image_url" example:"https://example.com/photo.jpg"`
}

func (adult *Adult) BeforeCreate(tx *gorm.DB) (err error) {
	adult.ID = uuid.New()
	return
}

type Child struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id" swaggertype:"string" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name      string    `json:"name" example:"Jane"`
	LastName  string    `json:"last_name" example:"Doe"`
	BirthYear int       `json:"birth_year" example:"2015"`
	ImageURL  string    `json:"image_url" example:"https://example.com/child_photo.jpg"`
}

func (child *Child) BeforeCreate(tx *gorm.DB) (err error) {
	child.ID = uuid.New()
	return
}

func InitDB() {
	var err error
	dsn := "host=localhost user=devuser password=devpassword123 dbname=intro_microservicios port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&Adult{}, &Child{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database connected and migrated successfully")
}

func main() {
	InitDB()

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
// @Success 200 {array} Adult
// @Router /Adults [get]
func GetAdults(c *gin.Context) {
	var adults []Adult
	if err := DB.Find(&adults).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adults)
}

// GetChildren godoc
// @Summary Get all children
// @Description Get a list of children
// @Tags Children
// @Produce  json
// @Success 200 {array} Child
// @Router /Children [get]
func GetChildren(c *gin.Context) {
	var children []Child
	if err := DB.Find(&children).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, children)
}

// GetAdultById godoc
// @Summary Get adult by ID
// @Description Get a specific adult by ID
// @Tags Adults
// @Produce  json
// @Param id path string true "Adult ID"
// @Success 200 {object} Adult
// @Router /Adults/{id} [get]
func GetAdultById(c *gin.Context) {
	id := c.Param("id")
	var adult Adult
	if err := DB.First(&adult, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Adult not found"})
		return
	}
	c.JSON(http.StatusOK, adult)
}

// GetChildById godoc
// @Summary Get child by ID
// @Description Get a specific child by ID
// @Tags Children
// @Produce  json
// @Param id path string true "Child ID"
// @Success 200 {object} Child
// @Router /Children/{id} [get]
func GetChildById(c *gin.Context) {
	id := c.Param("id")
	var child Child
	if err := DB.First(&child, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Child not found"})
		return
	}
	c.JSON(http.StatusOK, child)
}

// AddAdult godoc
// @Summary Add a new adult
// @Description Add a new adult
// @Tags Adults
// @Accept  json
// @Produce  json
// @Param adult body Adult true "Adult Object"
// @Success 200 {object} Response
// @Router /Add/Adults [post]
func AddAdult(c *gin.Context) {
	var adult Adult
	if err := c.ShouldBindJSON(&adult); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&adult).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Message: "Adult added successfully"})
}

// AddChild godoc
// @Summary Add a new child
// @Description Add a new child
// @Tags Children
// @Accept  json
// @Produce  json
// @Param child body Child true "Child Object"
// @Success 200 {object} Response
// @Router /Add/Children [post]
func AddChild(c *gin.Context) {
	var child Child
	if err := c.ShouldBindJSON(&child); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&child).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Message: "Child added successfully"})
}
