package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Place struct {
	Id          uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Description string         `gorm:"type:text;not null" json:"description"`
	Longitude   float64        `gorm:"type:decimal(10,8);not null" json:"longitude"`
	Latitude    float64        `gorm:"type:decimal(10,8);not null" json:"latitude"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdateAt    time.Time      `json:"update_at"`
	DeleteAt    gorm.DeletedAt `json:"delete_at"`
}
