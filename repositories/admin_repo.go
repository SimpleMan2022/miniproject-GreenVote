package repositories

import (
	"evoting/entities"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByEmail(email string) (*entities.Admin, error)
	SaveRefreshToken(user *entities.Admin) error
	GetUserByRefreshToken(token string) (*entities.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) FindByEmail(email string) (*entities.Admin, error) {
	var admin *entities.Admin
	if err := r.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *adminRepository) SaveRefreshToken(admin *entities.Admin) error {
	err := r.db.Save(&admin).Error
	return err
}

func (r *adminRepository) GetUserByRefreshToken(token string) (*entities.Admin, error) {
	var admin *entities.Admin
	err := r.db.Where("refresh_token = ?", token).First(&admin).Error
	return admin, err
}
