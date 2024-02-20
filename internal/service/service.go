package service

import (
	"github.com/mystpen/weather-api/internal/repository"
	"github.com/mystpen/weather-api/internal/service/weather"
)

type Service struct{
	WeatherService weather.WeatherServiceImplement
}

func NewService(repo *repository.Repository) *Service{
	return &Service{
		WeatherService: weather.NewWeatherService(repo.WeatherRepo),
	}
}