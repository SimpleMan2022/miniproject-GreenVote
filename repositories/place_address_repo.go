package repositories

import (
	"evoting/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceAddressRepository interface {
	FindById(id uuid.UUID) (*entities.PlaceAddress, error)
	Update(address *entities.PlaceAddress) (*entities.PlaceAddress, error)
	Delete(address *entities.PlaceAddress) error
}

type placeAddressRepository struct {
	db *gorm.DB
}

func NewPlaceAddressRepository(db *gorm.DB) *placeAddressRepository {
	return &placeAddressRepository{db}
}

func (r *placeAddressRepository) FindById(id uuid.UUID) (*entities.PlaceAddress, error) {
	var address entities.PlaceAddress
	if err := r.db.Where("id = ?", id).First(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil
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
