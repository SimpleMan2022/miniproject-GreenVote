package config

import (
	mysql2 "evoting/drivers/mysql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func LoadDb() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%s",
		ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_NAME, "utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&mysql2.Address{}, mysql2.User{})
	//if err := seeders.DBSeed(db); err != nil {
	//	panic(err)
	//}
	DB = db
}
