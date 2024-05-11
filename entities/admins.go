package entities

import (
	"github.com/google/uuid"
	"time"
)

type Admin struct {
	Id           uuid.UUID
	Email        string
	Username     string
	Password     string
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
