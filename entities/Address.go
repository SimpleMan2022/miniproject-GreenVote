package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	Id          uuid.UUID
	Province    string
	City        string
	SubDistrict string
	StreetName  string
	ZipCode     string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   gorm.DeletedAt
}
