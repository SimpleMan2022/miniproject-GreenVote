package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id           uuid.UUID
	Email        string
	Fullname     string
	Password     string
	AddressId    *uuid.UUID
	Address      *Address
	Image        *string
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
