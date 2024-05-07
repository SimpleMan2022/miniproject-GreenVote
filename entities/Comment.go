package entities

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	PlaceId   uuid.UUID
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
