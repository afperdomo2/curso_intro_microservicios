package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"afperdomo2/go/microservicios/services/add-member/kafka"
	"afperdomo2/go/microservicios/services/add-member/models"

	"github.com/gin-gonic/gin"
)

// AddMemberHandler encapsula las dependencias del handler
type AddMemberHandler struct {
	kafkaProducer *kafka.Producer
}

// NewAddMemberHandler crea una nueva instancia del handler con sus dependencias
func NewAddMemberHandler(kafkaProducer *kafka.Producer) *AddMemberHandler {
	return &AddMemberHandler{
		kafkaProducer: kafkaProducer,
	}
}

// AddMember maneja la solicitud POST para agregar un nuevo miembro
func (h *AddMemberHandler) AddMember(c *gin.Context) {
	var member models.Member

	// Validar el JSON recibido
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Crear el mensaje que se enviará a Kafka
	message := models.MemberMessage{
		Name:      member.Name,
		LastName:  member.LastName,
		BirthYear: member.BirthYear,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// Enviar mensaje a Kafka
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	messageKey := fmt.Sprintf("member-%s-%s", member.Name, member.LastName)
	if err := h.kafkaProducer.SendMessage(ctx, messageKey, message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to send message to Kafka",
			"details": err.Error(),
		})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Member %s %s added successfully", member.Name, member.LastName),
		"data":    message,
	})
}
