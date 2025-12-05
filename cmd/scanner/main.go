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
}
