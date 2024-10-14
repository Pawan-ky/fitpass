package database

import (
	"fitpass/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var Instance *gorm.DB

func Init(){
	//  initialize db
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("couldn't connect to database")
	}
	Instance = db
}

func Migrate() {
	err := Instance.AutoMigrate(&models.User{}, &models.Subscription{})
	if err != nil {
		panic(err.Error())
	}
}