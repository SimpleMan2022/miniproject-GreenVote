package mysql

import (
	"github.com/google/uuid"
	"time"
)

type Admin struct {
	Id           uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Email        string    `gorm:"type:varchar(255);not null" json:"email"`
	Username     string    `gorm:"type:varchar(255);not null" json:"username"`
	Password     string    `gorm:"type:varchar(255);not null" json:"password"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
