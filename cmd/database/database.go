package database

import (
	"awesomeProject/cmd/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func migrateDb() {
	DB.AutoMigrate(&models.User{})
}

func createDb() {
	db, err := gorm.Open(sqlite.Open("root.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	DB = db
}

func InitDB() {
	createDb()
	migrateDb()
}
