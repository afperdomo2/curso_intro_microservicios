package repository

import (
	"log"

	"afperdomo2/go/microservicios/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AdultRepository encapsula la lógica de acceso a datos para adultos
type AdultRepository struct {
	db *gorm.DB
}

// NewAdultRepository crea una nueva instancia del repositorio
func NewAdultRepository(db *gorm.DB) *AdultRepository {
	return &AdultRepository{
		db: db,
	}
}

// SaveAdult guarda un adulto en la base de datos
func (r *AdultRepository) SaveAdult(name, lastName string, birthYear int, imageURL string) error {
	adult := models.Adult{
		ID:        uuid.New(),
		Name:      name,
		LastName:  lastName,
		BirthYear: birthYear,
		ImageURL:  imageURL,
	}

	if err := r.db.Create(&adult).Error; err != nil {
		log.Printf("[ERROR] ❌ Error guardando adulto en BD: %v", err)
		return err
	}

	log.Printf("✅ Adulto guardado exitosamente: %s %s (ID: %s)", adult.Name, adult.LastName, adult.ID)
	return nil
}
