package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "skc.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&entry{})

	DB = database
}
