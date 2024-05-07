package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id        uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	UserID    uuid.UUID      `gorm:"type:varchar(191);index" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	PlaceId   uuid.UUID      `gorm:"type:varchar(191);index" json:"place_id"`
	Place     Place          `gorm:"foreignKey:PlaceId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Body      string         `json:"body"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
