package app

import (
	"log"

	"github.com/mystpen/weather-api/internal/config"
	"github.com/mystpen/weather-api/internal/handlers"
	"github.com/mystpen/weather-api/internal/repository"
	"github.com/mystpen/weather-api/internal/service"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func Run() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	client := NewMongoDB(config)
	defer CloseMongoDBConnection(*client)
	collection := NewMongoCollection(client, config, "weather")

	repo := repository.NewRepository(collection)
	service := service.NewService(repo)
	handler := handlers.NewHandler(service, config)

	log.Fatal(Server(handler.Routes()))
}
