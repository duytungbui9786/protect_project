package database

import (
	"github.com/TechMaster/golang/08Fiber/Basic/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("tung:12345678@/shop"), &gorm.Config{})
	if err != nil {
		panic("couldn't connect to the database")
	}

	DB = database

	database.AutoMigrate(&models.User{})
}
