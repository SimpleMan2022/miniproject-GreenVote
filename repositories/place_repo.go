package repositories

import (
	"evoting/entities"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type PlaceRepository interface {
	Create(place *entities.Place) (*entities.Place, error)
	CreateAddress(place *entities.PlaceAddress) (*entities.PlaceAddress, error)
	FindAll(page, limit int, sortBy, sortType, searchQuery string) (*[]entities.Place, *int64, error)
	FindByName(place string) (*entities.Place, error)
	FindById(id uuid.UUID) (*entities.Place, error)
	FindAddress(id uuid.UUID) (*entities.PlaceAddress, error)
	Update(place *entities.Place) (*entities.Place, error)
	UpdateAddress(place *entities.PlaceAddress) (*entities.PlaceAddress, error)
	Delete(place *entities.Place) error
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

func (r *placeRepository) FindAll(page, limit int, sortBy, sortType, searchQuery string) (*[]entities.Place, *int64, error) {
	var places []entities.Place
	var total int64
	offset := (page - 1) * limit
	db := r.db.Model(&entities.Place{})
	if sortBy != "" {
		db = db.Order(fmt.Sprintf("%s %s", sortBy, sortType))
	}
	if searchQuery != "" {
		db = db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(searchQuery)+"%")
	}

	if err := db.Debug().Preload("Address").Preload("Weather").
		Offset(offset).Limit(limit).
		Find(&places).
		Error; err != nil {
		return nil, nil, err
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	return &places, &total, nil
}

func (r *placeRepository) FindByName(name string) (*entities.Place, error) {
	var place *entities.Place
	if err := r.db.Where("name = ?", name).First(&place).Error; err != nil {
		return nil, err
	}
	return place, nil
}

func (r *placeRepository) Delete(place *entities.Place) error {
	if err := r.db.Unscoped().Delete(&place).Error; err != nil {
		return err
	}
	return nil
}

func (r *placeRepository) FindById(id uuid.UUID) (*entities.Place, error) {
	var place entities.Place
	if err := r.db.Where("id = ?", id).Preload("Address").Preload("Weather").
		First(&place).Error; err != nil {
		return nil, err
	}
	return &place, nil
}

func (r *placeRepository) FindAddress(id uuid.UUID) (*entities.PlaceAddress, error) {
	var placeAddress entities.PlaceAddress
	if err := r.db.Where("place_id = ?", id).First(&placeAddress).Error; err != nil {
		return nil, err
	}
	return &placeAddress, nil
}

func (r *placeRepository) Update(place *entities.Place) (*entities.Place, error) {
	if err := r.db.Save(&place).Error; err != nil {
		return nil, err
	}
	return place, nil
}

func (r *placeRepository) UpdateAddress(placeAddress *entities.PlaceAddress) (*entities.PlaceAddress, error) {
	if err := r.db.Save(&placeAddress).Error; err != nil {
		return nil, err
	}
	return placeAddress, nil
}
