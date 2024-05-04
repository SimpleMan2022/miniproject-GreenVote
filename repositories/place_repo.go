package repositories

import (
	"evoting/entities"
	"gorm.io/gorm"
)

type PlaceRepository interface {
	Create(place *entities.Place) (*entities.Place, error)
	CreateAddress(place *entities.PlaceAddress) (*entities.PlaceAddress, error)
}

type placeRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) *placeRepository {
	return &placeRepository{db}
}

func (r *placeRepository) Create(place *entities.Place) (*entities.Place, error) {
	if err := r.db.Create(&place).Error; err != nil {
		return nil, err
	}
	return place, nil
}

func (r *placeRepository) CreateAddress(placeAddress *entities.PlaceAddress) (*entities.PlaceAddress, error) {
	if err := r.db.Create(&placeAddress).Error; err != nil {
		return nil, err
	}
	return placeAddress, nil
}
