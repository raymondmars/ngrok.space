package database

import (
	"common-user/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	// github.com/mattn/go-sqlite3
	Db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	// Create the table from our struct.
	db.AutoMigrate(&models.User{})

}
