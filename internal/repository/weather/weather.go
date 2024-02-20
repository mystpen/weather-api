package weather

import (
	"context"
	"log"

	"github.com/mystpen/weather-api/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WeatherDB struct {
	Collection *mongo.Collection
}

func NewWeatherDB(collection *mongo.Collection) *WeatherDB {
	return &WeatherDB{Collection: collection}
}

type WeatherRepo interface {
	PutData(weatherData *model.WeatherData) error
	GetData(location string) (model.WeatherData, error)
}

func (wdb *WeatherDB) PutData(weatherData *model.WeatherData) error {
	filter := bson.M{"location": weatherData.Location}

	update := bson.M{
		"$set": bson.M{
			"description": weatherData.Description,
			"temperature": weatherData.Temperature,
		},
	}

	opts := options.Update().SetUpsert(true)

	result, err := wdb.Collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Matched %v document and modified %v document.\n", result.MatchedCount, result.ModifiedCount)
	return nil
}

func (wdb *WeatherDB) GetData(location string) (model.WeatherData, error){
	filter := bson.M{"location": location}

	var weatherData model.WeatherData
	err := wdb.Collection.FindOne(context.TODO(), filter).Decode(&weatherData)
	if err != nil {
		log.Println(err)
		return weatherData ,err
	}
	return weatherData, err
}