package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Email        string         `gorm:"type:varchar(255);not null" json:"email"`
	Fullname     string         `gorm:"type:varchar(255);not null" json:"fullname"`
	Password     string         `gorm:"type:varchar(255);not null" json:"password"`
	Image        *string        `gorm:"type:varchar(255);null" json:"image"`
	RefreshToken string         `json:"refresh_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
