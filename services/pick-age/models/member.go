package models

// Member representa los datos de un miembro recibido desde Kafka
type Member struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	BirthYear int    `json:"birth_year"`
	Timestamp string `json:"timestamp"`
}
