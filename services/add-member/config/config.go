package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	KafkaBroker string
}

func LoadConfig() *Config {
	// Cargar variables de entorno desde archivos .env
	_ = godotenv.Load()          // Busca en ./
	_ = godotenv.Load("../.env") // Busca en ../

	cfg := &Config{
		KafkaBroker: getEnv("KAFKA_BROKER"),
	}

	return cfg
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("[ERROR] ❌ La variable de entorno '%s' es obligatoria pero no está definida", key)
	}
	return value
}
