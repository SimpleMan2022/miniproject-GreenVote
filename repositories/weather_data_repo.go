package repositories

import (
	"evoting/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WeatherDataRepository interface {
	FindPlace(placeId uuid.UUID) (*entities.Place, error)
	FindByPlaceId(placeId uuid.UUID) (*entities.WeatherData, error)
	Create(data *entities.WeatherData) (*entities.WeatherData, error)
	Update(data *entities.WeatherData) (*entities.WeatherData, error)
	Delete(data *entities.WeatherData) error
}

type weatherDataRepository struct {
	db *gorm.DB
}

func NewWeatherDataRepository(db *gorm.DB) *weatherDataRepository {
	return &weatherDataRepository{db}
}

func (r *weatherDataRepository) FindPlace(placeId uuid.UUID) (*entities.Place, error) {
	var place entities.Place
	if err := r.db.Where("id = ?", placeId).First(&place).Error; err != nil {
		return nil, err
	}
	return &place, nil
}

func (r *weatherDataRepository) FindByPlaceId(placeId uuid.UUID) (*entities.WeatherData, error) {
	var weather entities.WeatherData
	if err := r.db.Where("place_id = ?", placeId).First(&weather).Error; err != nil {
		return nil, err
	}
	return &weather, nil
}

func (r *weatherDataRepository) Create(data *entities.WeatherData) (*entities.WeatherData, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *weatherDataRepository) Update(data *entities.WeatherData) (*entities.WeatherData, error) {
	if err := r.db.Save(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *weatherDataRepository) Delete(data *entities.WeatherData) error {
	if err := r.db.Unscoped().Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
