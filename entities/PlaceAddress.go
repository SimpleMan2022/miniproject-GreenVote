package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type PlaceAddress struct {
	Id          uuid.UUID
	PlaceId     uuid.UUID
	Province    string
	City        string
	SubDistrict string
	StreetName  string
	ZipCode     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
