package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"afperdomo2/go/microservicios/services/add-adult/repository"

	"github.com/segmentio/kafka-go"
)

// ClassifiedAdult representa un adulto clasificado recibido desde Kafka
type ClassifiedAdult struct {
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	BirthYear   int       `json:"birth_year"`
	ImageURL    string    `json:"image_url"`
	Age         int       `json:"age"`
	PublishedAt time.Time `json:"published_at"`
}

// Consumer es responsable de leer adultos clasificados y guardarlos en BD
type Consumer struct {
	reader     *kafka.Reader
	repository *repository.AdultRepository
}

// NewConsumer crea una nueva instancia del consumidor de Kafka
func NewConsumer(topic string, brokerAddr string, repo *repository.AdultRepository) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{brokerAddr},
		Topic:          topic,
		GroupID:        "add-adult-service",
		StartOffset:    kafka.LastOffset,
		CommitInterval: time.Second,
		MaxBytes:       10e6,
	})

	log.Printf("🔊 Kafka Consumer inicializado para topic '%s' en %s", topic, brokerAddr)
	return &Consumer{
		reader:     reader,
		repository: repo,
	}
}

// Start inicia la escucha de mensajes del tema (bloqueante)
func (c *Consumer) Start(ctx context.Context) error {
	log.Println("👂 Consumer escuchando mensajes de adultos clasificados...")

	for {
		// Revisar si el contexto fue cancelado
		select {
		case <-ctx.Done():
			log.Println("⛔ Contexto del consumer cancelado")
			return ctx.Err()
		default:
		}

		// Leer mensaje
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("[ERROR] ❌ Error al leer mensaje: %v", err)
			continue
		}

		// Deserializar el mensaje
		var classifiedAdult ClassifiedAdult
		if err := json.Unmarshal(msg.Value, &classifiedAdult); err != nil {
			log.Printf("[ERROR] ❌ Error al deserializar mensaje: %v", err)
			continue
		}

		// Procesar el adulto (guardar en BD)
		c.processAdult(classifiedAdult)
	}
}

// processAdult guarda el adulto clasificado en la base de datos
func (c *Consumer) processAdult(classified ClassifiedAdult) {
	if err := c.repository.SaveAdult(classified.Name, classified.LastName, classified.BirthYear, classified.ImageURL); err != nil {
		log.Printf("[ERROR] ❌ Error procesando adulto %s %s: %v",
			classified.Name, classified.LastName, err)
		return
	}

	log.Printf("👤 Adulto procesado exitosamente: %s %s (edad: %d años)",
		classified.Name, classified.LastName, classified.Age)
}

// Close cierra la conexión con Kafka
func (c *Consumer) Close() error {
	if err := c.reader.Close(); err != nil {
		return err
	}
	log.Println("🔌 Kafka Consumer cerrado")
	return nil
}
