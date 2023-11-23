package entity

import (
	"gorm.io/gorm"
	
	"time"
)

type Member struct {
	gorm.Model
	FirstName string
	LastName string
	Tel int
	Email string `gorm:"uniqueIndex"`
	Password string
	Point float32
}
type Admin struct {
	gorm.Model
	Email string `gorm:"uniqueIndex"`
	Password string 
}
type Room struct {
	gorm.Model
	Roomtype_id int //fk
	Room_name string
	Start_time  time.Time
	End_time time.Time
	room_status_id int //fk
}

type Room_type struct {
	gorm.Model
	Name string
	Capacity int 
}

type Room_status struct{
	gorm.Model
	Status_name string 
}
type Room_booking struct {
	gorm.Model
	Member_id int  //fk
	Room_id int //fk
	Start_time  time.Time
	End_time time.Time
}

type Rewards_history struct {
	gorm.Model 
	Rewardshis_date time.Time
	Member_id int //fk
	Rewards_id int //fk
}
type Rewards struct {
	gorm.Model
	Rewards_cat_id int  // fk
	Rewards_picture string 
	Rewards_title string 
	return_details string 
	Point_to_paid int 
	Rewards_stock int 
}

type Point_history struct {
	gorm.Model 
	Point int 
	Remark string
	Book_borrow_id int //fk
	Room_booking_id int //fk
	Member_id int //fk
	Donate_book_id int  //fk
}
type Donate_book struct {
	gorm.Model 
	DonateBook_title string
	Donate_point int 
	Cat_id int //fk
	Member_id int //fk	
}
type Book_status struct {
	gorm.Model
	Status_name string //fk
}
type Book_request struct {
	gorm.Model
	BookRequest_title string 
	BookRequest_reason string
	Book_request_status_id int //fk
	Member_id int //fk
	Cat_id int //fk 
}
type Catagory struct {
	gorm.Model 
	Cat_name string 
}
type Book struct {
	gorm.Model 
	Book_title string 
	QuantityInStock int 
	QuantityInRental int 
	Image_src string 
	TotalBook int 
}
type Borrow_status struct {
	gorm.Model
	Name string 
}
type Borrow_book struct {
	gorm.Model 
	Member_id int //fk
	Book_id int //fk
	Start_time time.Time
	End_time time.Time 
	Totalbook int 
	Borrow_status string //fk
}
type Review struct {
	gorm.Model
	Rate string
	Comment string 
	Book_id string //fk
	Member_id //fk 
}