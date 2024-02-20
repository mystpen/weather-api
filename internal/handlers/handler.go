package handlers

import "github.com/mystpen/weather-api/internal/service"

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
