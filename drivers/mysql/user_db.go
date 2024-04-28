package mysql

import "time"

type User struct {
	Id        int       `gorm:"primaryKey;not null" json:"id"`
	Email     string    `gorm:"type:varchar(255);not null" json:"email"`
	Fullname  string    `gorm:"type:varchar(255);not null" json:"fullname"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	Address   *Address  `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
