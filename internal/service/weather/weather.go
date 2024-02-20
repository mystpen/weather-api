package weather

import "github.com/mystpen/weather-api/internal/repository/weather"

type WeatherService struct{
	repo weather.WeatherRepo
}

func NewWeatherService(repo weather.WeatherRepo) *WeatherService{
	return &WeatherService{repo: repo}
}

type WeatherServiceImplement interface{

}