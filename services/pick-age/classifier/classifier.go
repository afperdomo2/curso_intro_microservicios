package classifier

import (
	"time"

	"afperdomo2/go/microservicios/services/pick-age/models"
)

// ClassificationType define los tipos de clasificación posibles
type ClassificationType string

const (
	Adult ClassificationType = "adult"
	Child ClassificationType = "child"
)

// MemberClassification contiene el resultado de clasificar a un miembro
type MemberClassification struct {
	Member models.Member
	Type   ClassificationType
	Age    int
}

// Classifier es responsable de clasificar miembros según su edad
type Classifier struct{}

// NewClassifier crea una nueva instancia del clasificador
func NewClassifier() *Classifier {
	return &Classifier{}
}

// Classify determina si un miembro es adulto o menor basado en su año de nacimiento
func (c *Classifier) Classify(member models.Member) *MemberClassification {
	currentYear := getCurrentYear()
	age := currentYear - member.BirthYear

	var classificationType ClassificationType
	if age >= 18 {
		classificationType = Adult
	} else {
		classificationType = Child
	}

	return &MemberClassification{
		Member: member,
		Type:   classificationType,
		Age:    age,
	}
}

// getCurrentYear devuelve el año actual
func getCurrentYear() int {
	return time.Now().Year()
}
