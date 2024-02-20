package model

type WeatherData struct {
	Location    string `bson:"location"`
	Temperature int    `bson:"temperature"`
	Description string `bson:"description"`
}
