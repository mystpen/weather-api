package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mystpen/weather-api/internal/config"
	"github.com/mystpen/weather-api/internal/model"
)

func GetWeatherInfo(config *config.Config, location string) (*model.WeatherData, error) {
	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", location, config.Token)

	var data model.WeatherData
	response, err := http.Get(apiUrl)
	if err != nil {
		return &data, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return &data, fmt.Errorf("request failed")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return &data, err
	}

	// Unmarshalling
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return &data, err
	}

	var currTemp float64
	var description string
	// Getting current temp
	if main, ok := result["main"].(map[string]interface{}); ok {
		if temp, ok := main["temp"].(float64); ok {
			currTemp = temp - 273.15
		}
	}
	// Getting weather description
	if weather, ok := result["weather"].([]interface{}); ok {
		if len(weather) > 0 {
			if description, ok = weather[0].(map[string]interface{})["description"].(string); ok {
				data.Description = description
			}
		}
	}

	data = model.WeatherData{
		Location:    location,
		Temperature: int(currTemp),
		Description: description,
	}

	return &data, nil
}
