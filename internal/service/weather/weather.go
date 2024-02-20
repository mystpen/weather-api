package weather

import (
	"github.com/mystpen/weather-api/internal/model"
	"github.com/mystpen/weather-api/internal/repository/weather"
)

type WeatherService struct{
	repo weather.WeatherRepo
}

func NewWeatherService(repo weather.WeatherRepo) *WeatherService{
	return &WeatherService{repo: repo}
}

type WeatherServiceImplement interface{
	PutData(weatherData *model.WeatherData) error
	GetData(location string) (model.WeatherData, error)
}

func (s *WeatherService) PutData(weatherData *model.WeatherData) error{
	return s.repo.PutData(weatherData)
}

func (s *WeatherService) GetData(location string) (model.WeatherData, error){
	return s.repo.GetData(location)
}