package entities

import (
	"github.com/google/uuid"
	"time"
)

type Votes struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	PlaceId   uuid.UUID
	CreatedAt time.Time
}
