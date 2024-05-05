package repositories

import (
	"evoting/entities"
	"gorm.io/gorm"
)

type PlaceAddressRepository interface {
	Update(address *entities.PlaceAddress) (*entities.PlaceAddress, error)
	Delete(address *entities.PlaceAddress) error
}

type placeAddressRepository struct {
	db *gorm.DB
}

func NewPlaceAddressRepository(db *gorm.DB) *placeAddressRepository {
	return &placeAddressRepository{db}
}

func (r *placeAddressRepository) Update(address *entities.PlaceAddress) (*entities.PlaceAddress, error) {
	if err := r.db.Save(&address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (r *placeAddressRepository) Delete(address *entities.PlaceAddress) error {
	if err := r.db.Delete(&address).Error; err != nil {
		return err
	}
	return nil
}
