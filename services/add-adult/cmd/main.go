package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"afperdomo2/go/microservicios/services/add-adult/config"
	"afperdomo2/go/microservicios/services/add-adult/kafka"
	"afperdomo2/go/microservicios/services/add-adult/repository"

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
	adultRepo := repository.NewAdultRepository(db)

	// Crear consumer para escuchar el topic de adultos clasificados
	kConsumer := kafka.NewConsumer("members.classification.fct.adult.validated", cfg.KafkaBroker, adultRepo)
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
	log.Println("🚀 Iniciando servicio AddAdult (Kafka Consumer)...")
	if err := kConsumer.Start(ctx); err != nil && err != context.Canceled {
		log.Fatalf("[ERROR] ❌ Error en consumer: %v", err)
	}

	log.Println("🛑 Servicio AddAdult detenido")
}
