package repository

import (
	"github.com/mystpen/weather-api/internal/repository/weather"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct{
	WeatherRepo weather.WeatherRepo
}

func NewRepository(collection *mongo.Collection) *Repository{
	return &Repository{
		WeatherRepo: weather.NewWeatherDB(collection),
	}
}