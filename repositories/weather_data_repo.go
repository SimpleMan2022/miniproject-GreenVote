package repositories

import (
	"evoting/entities"
	"gorm.io/gorm"
)

type WeatherDataRepository interface {
	FindByPlaceId(data *entities.WeatherData) (*entities.WeatherData, error)
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

func (r *weatherDataRepository) FindByPlaceId(data *entities.WeatherData) (*entities.WeatherData, error) {
	if err := r.db.Where("place_id = ?", data.PlaceId).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
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
	if err := r.db.Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
