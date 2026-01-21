package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// Producer encapsula la lógica para enviar mensajes a Kafka
type Producer struct {
	writer *kafka.Writer
}

// NewProducer crea una nueva instancia del productor de Kafka
func NewProducer(topic string, brokerAddr string) *Producer {

	// Crear el topic manualmente al inicio
	conn, err := kafka.Dial("tcp", brokerAddr)
	if err == nil {
		defer conn.Close()

		topicConfig := kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		}
		err = conn.CreateTopics(topicConfig)
		if err != nil {
			log.Printf("[WARN] ⚠️ Fallo al crear topic en Kafka (posiblemente ya existe): %v", err)
		}
	}

	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddr),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		MaxAttempts:  3,                // Para pruebas/local (10 en producción)
		RequiredAcks: kafka.RequireOne, // Para pruebas (kafka.RequireAll en producción)

		// Optimización de confiabilidad
		WriteTimeout: 5 * time.Second, // Tiempo de espera para escribir mensajes (10s en producción)
		ReadTimeout:  5 * time.Second, // Tiempo de espera para leer respuestas (10s en producción)

		AllowAutoTopicCreation: false, // Ya estamos creando el topic manualmente
	}

	log.Printf("Kafka Producer initialized for topic '%s' at %s", topic, brokerAddr)
	return &Producer{writer: writer}
}

// SendMessage envía un mensaje a Kafka
func (p *Producer) SendMessage(ctx context.Context, key string, value interface{}) error {
	// Serializar el mensaje a JSON
	messageBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error serializing message: %w", err)
	}

	// Crear el mensaje de Kafka
	msg := kafka.Message{
		Key:   []byte(key),
		Value: messageBytes,
	}

	// Enviar el mensaje
	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		return fmt.Errorf("error sending message to Kafka: %w", err)
	}

	log.Printf("Message sent to Kafka topic '%s': %s", p.writer.Topic, string(messageBytes))
	return nil
}

// Close cierra la conexión con Kafka
func (p *Producer) Close() error {
	if err := p.writer.Close(); err != nil {
		return fmt.Errorf("error closing Kafka writer: %w", err)
	}
	log.Println("Kafka Producer closed")
	return nil
}
