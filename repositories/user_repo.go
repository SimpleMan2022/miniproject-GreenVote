package repositories

import (
	"evoting/entities"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type UserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	FindById(id uuid.UUID) (*entities.User, error)
	FindAll(page, limit int, sortBy, sortType, searchQuery string) (*[]entities.User, *int64, error)
	FindSoftDelete(page, limit int, sortBy, sortType string) (*[]entities.User, *int64, error)
	Create(user *entities.User) (*entities.User, error)
	SaveRefreshToken(user *entities.User) error
	Update(user *entities.User) (*entities.User, error)
	Delete(user *entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*entities.User, error) {
	var user *entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindById(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("id = ?", id).Preload("Address").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll(page, limit int, sortBy, sortType, searchQuery string) (*[]entities.User, *int64, error) {
	var user []entities.User
	var total int64
	offset := (page - 1) * limit
	db := r.db.Model(&entities.User{})
	if sortBy != "" {
		db = db.Order(fmt.Sprintf("%s %s", sortBy, sortType))
	}
	if searchQuery != "" {
		db = db.Where("LOWER(fullname) LIKE ?", "%"+strings.ToLower(searchQuery)+"%")
	}

	if err := db.Preload("Address").
		Offset(offset).Limit(limit).
		Find(&user).
		Error; err != nil {
		return nil, nil, err
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	return &user, &total, nil
}

func (r *userRepository) FindSoftDelete(page, limit int, sortBy, sortType string) (*[]entities.User, *int64, error) {
	var user []entities.User
	var total int64
	offset := (page - 1) * limit
	db := r.db.Unscoped().Model(&entities.User{})
	if sortBy != "" {
		db = db.Order(fmt.Sprintf("%s %s", sortBy, sortType))
	}

	if err := db.Preload("Address").
		Where("deleted_at IS NOT NULL").
		Offset(offset).Limit(limit).
		Find(&user).
		Error; err != nil {
		return nil, nil, err
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	return &user, &total, nil
}

func (r *userRepository) Create(user *entities.User) (*entities.User, error) {

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) SaveRefreshToken(user *entities.User) error {
	err := r.db.Save(&user).Error
	return err
}

func (r *userRepository) Update(user *entities.User) (*entities.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}

func (r *userRepository) Delete(user *entities.User) error {
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
