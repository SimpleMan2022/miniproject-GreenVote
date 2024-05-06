package dto

import "github.com/google/uuid"

type WeatherDataRequest struct {
	PlaceId     uuid.UUID `json:"place_id"`
	Temperature float64   `json:"temperature"`
	WindSpeed   float64   `json:"wind_speed"`
	Humadity    float64   `json:"humadity"`
	Summary     string    `json:"summary"`
}
