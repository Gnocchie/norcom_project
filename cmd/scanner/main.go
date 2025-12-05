package main

import (
	"log"

	"github.com/Gnocchie/norcom_project/internal/config"
	"github.com/Gnocchie/norcom_project/internal/messaging"
	"github.com/Gnocchie/norcom_project/internal/scanner"
)

func main() {
	cfg := config.Load()

	publisher, err := messaging.NewRabbitPublisher(cfg)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ publisher: %v", err)
	}
	defer publisher.Close()

	log.Println("Starting file scan in:", cfg.RootPath)
	err = scanner.Walk(cfg.RootPath, publisher)
	if err != nil {
		log.Fatalf("File scanning error: %v", err)
	}

	log.Println("Scan complete.")
}
