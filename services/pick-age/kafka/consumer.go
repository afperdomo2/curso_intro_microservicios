package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"afperdomo2/go/microservicios/services/pick-age/models"

	"github.com/segmentio/kafka-go"
)

// Consumer encapsula la lógica para recibir y procesar mensajes de Kafka
type Consumer struct {
	reader *kafka.Reader
}

// NewConsumer crea una nueva instancia del consumidor de Kafka
func NewConsumer(topic string, brokerAddr string) *Consumer {
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
		reader: reader,
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

		// Procesar el miembro (calcular edad y loguear)
		c.processMember(member)
	}
}

// processMember analiza la edad del miembro y loguea si es adulto o menor
func (c *Consumer) processMember(member models.Member) {
	currentYear := time.Now().Year()
	age := currentYear - member.BirthYear

	if age >= 18 {
		log.Printf("👤 ADULTO: %s %s - Nacido en %d (edad: %d años)",
			member.Name, member.LastName, member.BirthYear, age)
	} else {
		log.Printf("👶 MENOR: %s %s - Nacido en %d (edad: %d años)",
			member.Name, member.LastName, member.BirthYear, age)
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
