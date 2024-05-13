package dto

import (
	"evoting/entities"
	"github.com/google/uuid"
	"time"
)

type WeatherDataRequest struct {
	PlaceId     uuid.UUID `json:"place_id"`
	Temperature int       `json:"temperature"`
	WindSpeed   float64   `json:"wind_speed"`
	Humadity    float64   `json:"humadity"`
	Condition   string    `json:"condition"`
}

type WeatherDataResponse struct {
	Place       PlaceRequest `json:"place"`
	Temperature int          `json:"temperature"`
	WindSpeed   float64      `json:"wind_speed"`
	Humadity    float64      `json:"humadity"`
	Condition   string       `json:"condition"`
	RecordedAt  time.Time    `json:"recorded_at"`
}

type WeatherDataPlace struct {
	Temperature int       `json:"temperature"`
	WindSpeed   float64   `json:"wind_speed"`
	Humadity    float64   `json:"humadity"`
	Condition   string    `json:"condition"`
	RecordedAt  time.Time `json:"recorded_at"`
}

func ToWeatherDataResponse(data *entities.WeatherData, place *entities.Place) *WeatherDataResponse {
	return &WeatherDataResponse{
		Place: PlaceRequest{
			Id:          place.Id,
			Name:        place.Name,
			Description: place.Description,
			Longitude:   place.Longitude,
			Latitude:    place.Latitude,
			MapImage:    place.MapImage,
		},
		Temperature: data.Temperature,
		WindSpeed:   data.WindSpeed,
		Humadity:    data.Humadity,
		Condition:   data.Condition,
		RecordedAt:  data.RecordedAt,
	}
}
