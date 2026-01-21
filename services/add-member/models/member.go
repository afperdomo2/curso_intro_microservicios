package models

// Member representa los datos de un miembro que será enviado a Kafka
type Member struct {
	Name      string `json:"name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	BirthYear int    `json:"birth_year" binding:"required"`
	ImageURL  string `json:"image_url"`
}

// MemberMessage representa el mensaje que se enviará a Kafka con información adicional
type MemberMessage struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	BirthYear int    `json:"birth_year"`
	ImageURL  string `json:"image_url"`
	Timestamp string `json:"timestamp"`
}
