package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"afperdomo2/go/microservicios/services/pick-age/classifier"

	"github.com/segmentio/kafka-go"
)

// PublishedMember es la estructura que se envía a Kafka tras la clasificación
type PublishedMember struct {
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	BirthYear   int       `json:"birth_year"`
	Age         int       `json:"age"`
	PublishedAt time.Time `json:"published_at"`
}

// Producer es responsable de publicar mensajes clasificados a Kafka
type Producer struct {
	writer *kafka.Writer
}

// NewProducer crea una nueva instancia del productor de Kafka
func NewProducer(brokerAddr string) *Producer {
	topics := []string{
		"members.classification.fct.adult.validated",
		"members.classification.fct.child.validated",
	}
	// Crear los topics manualmente al inicio
	for _, topic := range topics {
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
				log.Printf("[WARN] ⚠️ Fallo al crear topic '%s' en Kafka (posiblemente ya existe): %v", topic, err)
			}
		}
	}

	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokerAddr),
		Balancer: &kafka.LeastBytes{},
	}

	log.Printf("📨 Kafka Producer inicializado en %s", brokerAddr)
	return &Producer{
		writer: writer,
	}
}

// PublishClassification publica un miembro clasificado al topic correspondiente
func (p *Producer) PublishClassification(ctx context.Context, classification *classifier.MemberClassification) error {
	// Determinar el topic según el tipo de clasificación
	topic := p.getTopicByClassification(classification.Type)

	// Crear el mensaje a publicar
	publishedMember := PublishedMember{
		Name:        classification.Member.Name,
		LastName:    classification.Member.LastName,
		BirthYear:   classification.Member.BirthYear,
		Age:         classification.Age,
		PublishedAt: time.Now(),
	}

	// Serializar el mensaje
	value, err := json.Marshal(publishedMember)
	if err != nil {
		log.Printf("[ERROR] ❌ Error serializando clasificación: %v", err)
		return err
	}

	// Enviar el mensaje a Kafka
	message := kafka.Message{
		Topic: topic,
		Key:   []byte(classification.Member.Name + ":" + classification.Member.LastName),
		Value: value,
	}

	err = p.writer.WriteMessages(ctx, message)
	if err != nil {
		log.Printf("[ERROR] ❌ Error publicando a topic '%s': %v", topic, err)
		return err
	}

	// Log de éxito
	emoji := "👤"
	if classification.Type == classifier.Child {
		emoji = "👶"
	}
	log.Printf("%s %s publicado a topic '%s' (edad: %d años)", emoji,
		classification.Member.Name+" "+classification.Member.LastName, topic, classification.Age)

	return nil
}

// getTopicByClassification retorna el topic según el tipo de clasificación
func (p *Producer) getTopicByClassification(classType classifier.ClassificationType) string {
	if classType == classifier.Adult {
		return "members.classification.fct.adult.validated"
	}
	return "members.classification.fct.child.validated"
}

// Close cierra la conexión del productor
func (p *Producer) Close() error {
	if err := p.writer.Close(); err != nil {
		return err
	}
	log.Println("🔌 Kafka Producer cerrado")
	return nil
}
