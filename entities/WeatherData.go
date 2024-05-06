package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type WeatherData struct {
	Id          uuid.UUID
	PlaceId     uuid.UUID
	Temperature float64
	WindSpeed   float64
	Humadity    float64
	Summary     string
	RecordedAt  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
