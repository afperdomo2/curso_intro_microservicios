package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"afperdomo2/go/microservicios/services/pick-age/classifier"
	"afperdomo2/go/microservicios/services/pick-age/models"

	"github.com/segmentio/kafka-go"
)

// Consumer es responsable únicamente de leer mensajes de Kafka
type Consumer struct {
	reader     *kafka.Reader
	classifier *classifier.Classifier
	publisher  *Producer
}

// NewConsumer crea una nueva instancia del consumidor de Kafka
func NewConsumer(topic string, brokerAddr string, publisher *Producer) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{brokerAddr},
		Topic:          topic,
		GroupID:        "pick-age-service",
		StartOffset:    kafka.LastOffset,
		CommitInterval: time.Second,
		MaxBytes:       10e6,
	})

	log.Printf("🔊 Kafka Consumer inicializado para topic '%s' en %s", topic, brokerAddr)
	return &Consumer{
		reader:     reader,
		classifier: classifier.NewClassifier(),
		publisher:  publisher,
	}
}

// Start inicia la escucha de mensajes del tema (bloqueante)
func (c *Consumer) Start(ctx context.Context) error {
	log.Println("👂 Consumer escuchando mensajes...")

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
		var member models.Member
		if err := json.Unmarshal(msg.Value, &member); err != nil {
			log.Printf("[ERROR] ❌ Error al deserializar mensaje: %v", err)
			continue
		}

		// Procesar el miembro (clasificar y publicar)
		c.processMember(ctx, member)
	}
}

// processMember es responsable de orquestar la clasificación y publicación
func (c *Consumer) processMember(ctx context.Context, member models.Member) {
	// Clasificar el miembro
	classification := c.classifier.Classify(member)

	// Publicar la clasificación al topic correspondiente
	if err := c.publisher.PublishClassification(ctx, classification); err != nil {
		log.Printf("[ERROR] ❌ Error publicando clasificación para %s %s: %v",
			member.Name, member.LastName, err)
	}
}

// Close cierra la conexión con Kafka
func (c *Consumer) Close() error {
	if err := c.reader.Close(); err != nil {
		return err
	}
	log.Println("🔌 Kafka Consumer cerrado")
	return nil
}
