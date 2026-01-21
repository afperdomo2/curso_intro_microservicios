package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"afperdomo2/go/microservicios/services/pick-age/config"
	"afperdomo2/go/microservicios/services/pick-age/kafka"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Crear kConsumer para escuchar el topic
	kConsumer := kafka.NewConsumer("members.registration.fct.member.received", cfg.KafkaBroker)
	defer kConsumer.Close()

	// Crear contexto para cancelar el consumer gracefully
	ctx, cancel := context.WithCancel(context.Background())

	// Manejar señales de cierre
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		log.Println("📍 Apagando gracefully...")
		cancel()
		kConsumer.Close()
		os.Exit(0)
	}()

	// Iniciar el consumer (bloqueante)
	log.Println("🚀 Iniciando servicio PickAge...")
	if err := kConsumer.Start(ctx); err != nil && err != context.Canceled {
		log.Fatalf("[ERROR] ❌ Error en consumer: %v", err)
	}

	log.Println("🛑 Servicio PickAge detenido")
}
