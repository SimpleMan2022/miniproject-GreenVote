package repositories

import (
	"evoting/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	Create(user *entities.User) (*entities.User, error)
	SaveRefreshToken(user *entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*entities.User, error) {
	var user *entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
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
