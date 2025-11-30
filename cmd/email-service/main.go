package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"aptiverse-email/internal/config"
	"aptiverse-email/internal/rabbitmq"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize RabbitMQ consumer
	consumer, err := rabbitmq.NewConsumer(cfg)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	// Start the consumer - THIS IS MISSING!
	err = consumer.Start()
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	log.Println("Email service started successfully")

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutdown signal received")

	consumer.Stop()
	log.Println("Email service stopped gracefully")
}