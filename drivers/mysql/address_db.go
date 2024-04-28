package mysql

import "github.com/google/uuid"

type Address struct {
	Id          uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Province    string    `gorm:"type:varchar(255);not null" json:"province"`
	City        string    `gorm:"type:varchar(255);not null" json:"city"`
	SubDistrict string    `gorm:"type:varchar(255);not null" json:"sub_district"`
	StreetName  string    `gorm:"type:varchar(255);not null" json:"street_name"`
	ZipCode     string    `gorm:"type:char(7);not null" json:"zip_code"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
