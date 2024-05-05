package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type PlaceAddress struct {
	ID          uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Place       Place          `gorm:"foreignKey:PlaceId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	PlaceId     uuid.UUID      `gorm:"type:varchar(191);index" json:"user_id"`
	Province    string         `gorm:"type:varchar(255);not null" json:"province"`
	City        string         `gorm:"type:varchar(255);not null" json:"city"`
	SubDistrict string         `gorm:"type:varchar(255);not null" json:"sub_district"`
	StreetName  string         `gorm:"type:varchar(255);not null" json:"street_name"`
	ZipCode     string         `gorm:"type:char(7);not null" json:"zip_code"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
