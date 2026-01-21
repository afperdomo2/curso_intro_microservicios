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

	// Crear el productor de Kafka
	kProducer := kafka.NewProducer("members.registration.fct.member.received", cfg.KafkaBroker)
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
		log.Println("Apagando gracefully...")
		kProducer.Close()
		os.Exit(0)
	}()

	// Iniciar el servidor
	log.Println("Servicio AddMember iniciando en :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Fallo al iniciar servidor: %v", err)
	}
}
