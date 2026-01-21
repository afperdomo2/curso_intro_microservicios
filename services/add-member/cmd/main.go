package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"afperdomo2/go/microservicios/services/add-member/config"
	"afperdomo2/go/microservicios/services/add-member/handlers"
	"afperdomo2/go/microservicios/services/add-member/kafka"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	// Crear el productor de Kafka para el topic "pickage"
	kProducer := kafka.NewProducer("pickage", cfg.KafkaBroker)
	defer kProducer.Close()

	// Crear el handler con sus dependencias inyectadas
	memberHandler := handlers.NewAddMemberHandler(kProducer)

	// Configurar Gin
	r := gin.Default()

	// Registrar la ruta POST
	r.POST("/Add/Member", memberHandler.AddMember)

	// Manejar señales de cierre graceful
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		log.Println("Shutting down gracefully...")
		kProducer.Close()
		os.Exit(0)
	}()

	// Iniciar el servidor
	log.Println("AddMember service starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
