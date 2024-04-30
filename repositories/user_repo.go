package repositories

import (
	"evoting/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	FindById(id uuid.UUID) (*entities.User, error)
	FindAll() (*[]entities.User, error)
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
	var user *entities.User
	if err := r.db.Where("id = ?", id).Preload("Address").First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindAll() (*[]entities.User, error) {
	var user *[]entities.User
	if err := r.db.Preload("Address").Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
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
