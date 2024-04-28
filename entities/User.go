package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	Email     string
	Fullname  string
	Password  string
	AddressId *Address
	Image     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
