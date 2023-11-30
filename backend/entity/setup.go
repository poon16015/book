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
	 &Room{},
	 &Admin{},
	 &Room_type{},
	 &Room_status{},
	 &Room_booking{},
	 &Rewards{},
	 &Rewards_history{},
	 &Point_history{},
	 &Donate_book{},
	 &BookStatus{},
	 &Book_request{},
	 &Catagory{},
	 &Book{},
	 &Borrow_status{},
	 &Borrow_book{},
	 &Review{},
	 &Rewards_catagory{},
	)
	db = database

	// db.Model(&Catagory{
	// 	Name: "Good",
	// })
  
	
  }
