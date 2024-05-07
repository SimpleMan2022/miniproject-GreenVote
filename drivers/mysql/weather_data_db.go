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
	Temperature int       `gorm:"type:int(11);not null" json:"temperature"`
	WindSpeed   float64   `gorm:"type:decimal(5,3);not null" json:"wind_speed"`
	Humadity    float64   `gorm:"type:decimal(5,3);not null" json:"humadity"`
	Condition   string    `gorm:"type:char(10);not null" json:"condition"`
	RecordedAt  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
