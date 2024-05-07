package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type WeatherData struct {
	Id          uuid.UUID
	PlaceId     uuid.UUID
	Temperature int
	WindSpeed   float64
	Humadity    float64
	Condition   string
	RecordedAt  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
