package mysql

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Email     string    `gorm:"type:varchar(255);not null" json:"email"`
	Fullname  string    `gorm:"type:varchar(255);not null" json:"fullname"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	AddressId uuid.UUID `gorm:"type:varchar(191);" json:"address_id"`
	Image     *string   `gorm:"type:varchar(255);null" json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}