package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

// Response represents a generic response
type Response struct {
	Message string `json:"message"`
}
