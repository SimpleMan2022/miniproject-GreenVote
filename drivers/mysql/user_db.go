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
	AddressId *int      `json:"address_id"`
	Image     *string   `gorm:"type:varchar(255);not null" json:"image"`
	Address   Address   `gorm:"foreignKey:AddressId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
