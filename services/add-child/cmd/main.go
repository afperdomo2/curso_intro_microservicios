package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"afperdomo2/go/microservicios/services/add-child/config"
	"afperdomo2/go/microservicios/services/add-child/kafka"
	"afperdomo2/go/microservicios/services/add-child/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Conectar a la base de datos
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[ERROR] ❌ Error conectando a la base de datos: %v", err)
	}

	log.Println("✅ Conectado a la base de datos")

	// Crear repositorio
	childRepo := repository.NewChildRepository(db)

	// Crear consumer para escuchar el topic de menores clasificados
	kConsumer := kafka.NewConsumer("members.classification.fct.child.validated", cfg.KafkaBroker, childRepo)
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
	log.Println("🚀 Iniciando servicio AddChild (Kafka Consumer)...")
	if err := kConsumer.Start(ctx); err != nil && err != context.Canceled {
		log.Fatalf("[ERROR] ❌ Error en consumer: %v", err)
	}

	log.Println("🛑 Servicio AddChild detenido")
}
