package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	KafkaBroker string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
}

func LoadConfig() *Config {
	// Cargar variables de entorno desde archivos .env
	_ = godotenv.Load()
	_ = godotenv.Load("../.env")

	cfg := &Config{
		KafkaBroker: getEnv("KAFKA_BROKER"),
		DBHost:      getEnvOrDefault("DB_HOST", "localhost"),
		DBPort:      getEnvOrDefault("DB_PORT", "5433"),
		DBUser:      getEnvOrDefault("DB_USER", "devuser"),
		DBPassword:  getEnvOrDefault("DB_PASSWORD", "devpassword123"),
		DBName:      getEnvOrDefault("DB_NAME", "intro_microservicios"),
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

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
