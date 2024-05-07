package entities

import (
	"github.com/google/uuid"
	"time"
)

type Place struct {
	Id          uuid.UUID
	Name        string
	Description string
	Longitude   float64
	Latitude    float64
	Address     PlaceAddress
	Weather     WeatherData
	MapImage    *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
