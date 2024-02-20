package handlers

import (
	"github.com/mystpen/weather-api/internal/config"
	"github.com/mystpen/weather-api/internal/service"
)

type Handler struct {
	service *service.Service
	config *config.Config
}

func NewHandler(service *service.Service, config *config.Config) *Handler {
	return &Handler{
		service: service,
		config: config,
	}
}
