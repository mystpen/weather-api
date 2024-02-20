package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mystpen/weather-api/internal/api"
)

func (h *Handler) weather(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/weather" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println("/////")
		return
	}
	switch r.Method {
	case http.MethodGet:
		location := r.URL.Query().Get("location")
		// get data from DB
		data, err := h.service.WeatherService.GetData(location)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// marshalling
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	case http.MethodPut:
		location := r.URL.Query().Get("location")
		// update or put weather data
		weatherData, err := api.GetWeatherInfo(h.config, location)
		if err != nil { 
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.service.WeatherService.PutData(weatherData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
