package weather

import "go.mongodb.org/mongo-driver/mongo"

type WeatherDB struct{
	Collection *mongo.Collection
}

func NewWeatherDB(collection *mongo.Collection) *WeatherDB{
	return &WeatherDB{Collection: collection}
}

type WeatherRepo interface{
	
}

