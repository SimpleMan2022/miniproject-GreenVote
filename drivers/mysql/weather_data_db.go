package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type WeatherData struct {
	Id          uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Place       Place     `gorm:"foreignKey:PlaceId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	PlaceId     uuid.UUID `gorm:"type:varchar(191);index" json:"place_id"`
	Temperature float64   `gorm:"type:decimal(11,8);not null" json:"temperature"`
	WindSpeed   float64   `gorm:"type:decimal(11,8);not null" json:"wind_speed"`
	Humadity    float64   `gorm:"type:decimal(11,8);not null" json:"humadity"`
	Summary     string    `gorm:"type:varchar(255);not null" json:"summary"`
	RecordedAt  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
