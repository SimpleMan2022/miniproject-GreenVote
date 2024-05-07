package mysql

import (
	"github.com/google/uuid"
	"time"
)

type Vote struct {
	Id        uuid.UUID `json:"id"`
	UserID    uuid.UUID `gorm:"type:varchar(191);index" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	PlaceId   uuid.UUID `gorm:"type:varchar(191);index" json:"place_id"`
	Place     Place     `gorm:"foreignKey:PlaceId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
}
