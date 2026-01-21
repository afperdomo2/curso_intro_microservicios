package repository

import (
	"log"

	"afperdomo2/go/microservicios/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ChildRepository encapsula la lógica de acceso a datos para menores
type ChildRepository struct {
	db *gorm.DB
}

// NewChildRepository crea una nueva instancia del repositorio
func NewChildRepository(db *gorm.DB) *ChildRepository {
	return &ChildRepository{
		db: db,
	}
}

// SaveChild guarda un menor en la base de datos
func (r *ChildRepository) SaveChild(name, lastName string, birthYear int, imageURL string) error {
	child := models.Child{
		ID:        uuid.New(),
		Name:      name,
		LastName:  lastName,
		BirthYear: birthYear,
		ImageURL:  imageURL,
	}

	if err := r.db.Create(&child).Error; err != nil {
		log.Printf("[ERROR] ❌ Error guardando menor en BD: %v", err)
		return err
	}

	log.Printf("✅ Menor guardado exitosamente: %s %s (ID: %s)", child.Name, child.LastName, child.ID)
	return nil
}
