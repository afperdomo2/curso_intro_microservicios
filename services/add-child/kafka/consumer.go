package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"afperdomo2/go/microservicios/services/add-child/repository"

	"github.com/segmentio/kafka-go"
)

// ClassifiedChild representa un menor clasificado recibido desde Kafka
type ClassifiedChild struct {
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	BirthYear   int       `json:"birth_year"`
	ImageURL    string    `json:"image_url"`
	Age         int       `json:"age"`
	ClassType   string    `json:"classification_type"`
	PublishedAt time.Time `json:"published_at"`
}

// Consumer es responsable de leer menores clasificados y guardarlos en BD
type Consumer struct {
	reader     *kafka.Reader
	repository *repository.ChildRepository
}

// NewConsumer crea una nueva instancia del consumidor de Kafka
func NewConsumer(topic string, brokerAddr string, repo *repository.ChildRepository) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{brokerAddr},
		Topic:          topic,
		GroupID:        "add-child-service",
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
	log.Println("👂 Consumer escuchando mensajes de menores clasificados...")

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
		var classifiedChild ClassifiedChild
		if err := json.Unmarshal(msg.Value, &classifiedChild); err != nil {
			log.Printf("[ERROR] ❌ Error al deserializar mensaje: %v", err)
			continue
		}

		// Procesar el menor (guardar en BD)
		c.processChild(classifiedChild)
	}
}

// processChild guarda el menor clasificado en la base de datos
func (c *Consumer) processChild(classified ClassifiedChild) {
	if err := c.repository.SaveChild(classified.Name, classified.LastName, classified.BirthYear, classified.ImageURL); err != nil {
		log.Printf("[ERROR] ❌ Error procesando menor %s %s: %v",
			classified.Name, classified.LastName, err)
		return
	}

	log.Printf("👶 Menor procesado exitosamente: %s %s (edad: %d años)",
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
