package seeders

import (
	"evoting/drivers/mysql/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Faker any
}

func RegisterSeeder(db *gorm.DB) []Seeder {
	var seeders []Seeder
	for i := 0; i < 20; i++ {
		seeders = append(seeders, Seeder{Faker: fakers.UserFaker(db)})
	}
	return seeders
}

func DBSeed(db *gorm.DB) error {
	seeders := RegisterSeeder(db)
	for _, seeder := range seeders {
		if err := db.Create(seeder.Faker).Error; err != nil {
			return err
		}
	}
	return nil
}
