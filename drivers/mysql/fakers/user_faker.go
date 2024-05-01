package fakers

import (
	"evoting/entities"
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func UserFaker(db *gorm.DB) *entities.User {
	return &entities.User{
		Id:           uuid.New(),
		Email:        faker.Email(),
		Fullname:     fmt.Sprintf("%s %s", faker.FirstName(), faker.LastName()),
		Address:      entities.Address{},
		Password:     "$2a$12$YKLHrrhcR/OSnb1QVmtEBuznBIHk3trQHGYNVq3QVaA4pgb7/v7Yi",
		Image:        nil,
		RefreshToken: "",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    gorm.DeletedAt{},
	}
}
