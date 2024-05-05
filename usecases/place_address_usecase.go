package usecases

import "github.com/google/uuid"

type PlaceAddressUsecase interface {
	Update(id uuid.UUID)
}
