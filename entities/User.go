package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id           uuid.UUID      `json:"id"`
	Email        string         `json:"email"`
	Fullname     string         `json:"fullname"`
	Password     string         `json:"password"`
	AddressId    *uuid.UUID     `json:"address_id"`
	Address      *Address       `json:"address"`
	Image        *string        `json:"image"`
	RefreshToken string         `json:"refresh_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
