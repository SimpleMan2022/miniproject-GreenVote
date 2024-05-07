package helpers

import (
	"encoding/json"
	"evoting/entities"
	"fmt"
	"github.com/spf13/viper"
	"math"
	"net/http"
)

type ResponseWeatherApi struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
}

func GenerateWeatherData(latitude, longitude float64) (*entities.WeatherData, error) {
	apiKey := viper.GetString("WEATHER_API_KEY")
	apiUrl := viper.GetString("WEATHER_API_URL")

	url := fmt.Sprintf(apiUrl, latitude, longitude, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ResponseWeatherApi
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	celcius := response.Main.Temp - 273.15

	weatherData := &entities.WeatherData{
		Temperature: int(math.Round(celcius)),
		WindSpeed:   response.Wind.Speed,
		Humadity:    response.Main.Humidity,
		Condition:   response.Weather[0].Main,
	}

	return weatherData, nil
}
