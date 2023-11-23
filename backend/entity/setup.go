package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("bookhouse.db"), &gorm.Config{})
	if err != nil {
	  panic("Failed to connect bookhouse database")
	}
  
	database.AutoMigrate(
	 &Member{},

	)
	db = database
  
	
  }
