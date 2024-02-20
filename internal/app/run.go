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
	handler := handlers.NewHandler(service)

	log.Fatal(Server(handler.Routes()))
	// ////
	// ash := Trainer{"Ash", 10, "Pallet Town"}
	// // misty := Trainer{"Misty", 10, "Cerulean City"}
	// // brock := Trainer{"Brock", 15, "Pewter City"}

	// insertResult, err := collection.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// filter := bson.D{{"name", "Ash"}}

	// // create a value into which the result can be decoded
	// var result Trainer

	// err = collection.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Found a single document: %+v\n", result)
}
