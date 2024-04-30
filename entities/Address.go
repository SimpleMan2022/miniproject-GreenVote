package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Address struct {
	Id          uuid.UUID      `json:"id"`
	Province    string         `json:"province"`
	City        string         `json:"city"`
	SubDistrict string         `json:"sub_district"`
	StreetName  string         `json:"street_name"`
	ZipCode     string         `json:"zip_code"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
