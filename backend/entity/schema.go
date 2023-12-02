package entity

import (
	"gorm.io/gorm"

	"time"
)

// type Prefixes struct {
// 	gorm.Model
// 	Name string
// }

// type Genders struct {
// 	gorm.Model
// 	Name string
// }

type Member struct {
	gorm.Model
	FirstName string
	LastName  string
	Tel       int
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Point     float32
	// GenderID  uint
	// PrefixID  uint
	Review          []Review          `gorm:"foreignKey:Member_id"`
	Borrow_book     []Borrow_book     `gorm:"foreignKey:Member_id"`
	Rewards_history []Rewards_history `gorm:"foreignKey:Member_id"`
	Donate_book     []Donate_book     `gorm:"foreignKey:Member_id"`
	Point_history   []Point_history   `gorm:"foreignKey:Member_id"`
	Room_booking    []Room_booking    `gorm:"foreignKey:Member_id"`
	Book_request    []Book_request    `gorm:"foreignKey:Member_id"`
}
type Admin struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string
}
type Room struct {
	gorm.Model

	Room_name      string
	Start_time     time.Time
	End_time       time.Time
	room_status_id int //fk

	RoomType_id *uint     //fk
	Room_type   Room_type `gorm:"references:id"`

	Room_status_id *uint       //fk
	Room_status    Room_status `gorm:"references:id"`

	Room_booking []Room_booking `gorm:"foreignKey:Room_id"`
}

type Room_type struct {
	gorm.Model
	Name     string
	Capacity int

	Room []Room `gorm:"foreignKey:RoomType_id"`
}

type Room_status struct {
	gorm.Model
	Status_name string

	Room []Room `gorm:"foreignKey:Room_status_id"`
}
type Room_booking struct {
	gorm.Model

	Start_time time.Time
	End_time   time.Time

	Member_id *uint  //fk
	Member    Member `gorm:"references:id"`

	Room_id *uint //fk
	Room    Room  `gorm:"references:id"`

	Point_history []Point_history `gorm:"foreignKey:Room_booking_id"`
}

type Rewards_history struct {
	gorm.Model
	Rewardshis_date time.Time

	Rewards_id *uint   //fk
	Rewards    Rewards `gorm:"references:id"`

	Member_id *uint  //fk
	Member    Member `gorm:"references:id"`
}
type Rewards struct {
	gorm.Model

	Rewards_picture string
	Rewards_title   string
	return_details  string
	Point_to_paid   int
	Rewards_stock   int

	Rewards_cat_id   *uint            // fk
	Rewards_catagory Rewards_catagory `gorm:"references:id"`

	Rewards_history []Rewards_history `gorm:"foreignKey:Rewards_id"`
}

type Rewards_catagory struct {
	gorm.Model
	Re_cat_name string

	Rewards []Rewards `gorm:"foreignKey:Rewards_cat_id"`
}

type Point_history struct {
	gorm.Model
	Point  int
	Remark string

	Member_id *uint  //fk
	Member    Member `gorm:"references:id"`

	Borrow_book_id *uint       //fk
	Borrow_book    Borrow_book `gorm:"references:id"`

	Room_booking_id *uint        //fk
	Room_booking    Room_booking `gorm:"references:id"`

	Donate_book_id *uint       //fk
	Donate_book    Donate_book `gorm:"references:id"`
}

type Donate_book struct {
	gorm.Model
	DonateBook_title string
	Donate_point     int

	Member_id uint   //fk
	Member    Member `gorm:"references:id"`

	Catagory_id uint     //fk
	Catagory    Catagory `gorm:"references:id"`

	Point_history []Point_history `gorm:"foreignKey:Donate_book_id"`
}

type BookStatus struct {
	gorm.Model
	StatusName string

	Book_request []Book_request `gorm:"foreignKey:BookStatus_id"`
}

type Book_request struct {
	gorm.Model
	BookRequest_title  string
	BookRequest_reason string

	Member_id *uint  //fk
	Member    Member `gorm:"references:id"`

	Catagory_id *uint    //fk
	Catagory    Catagory `gorm:"references:id"`

	BookStatus_id *uint       //fk
	BookStatus    BookStatus `gorm:"references:id"`
}

type Catagory struct {
	gorm.Model
	Name string

	Book         []Book         `gorm:"foreignKey:Catagory_id"`
	Book_request []Book_request `gorm:"foreignKey:Catagory_id"`
	Donate_book  []Donate_book  `gorm:"foreignKey:Catagory_id"`
}
type Book struct {
	gorm.Model
	Book_title       string
	QuantityInStock  int
	QuantityInRental int
	Image_src        string
	TotalBook        int

	Review []Review `gorm:"foreignKey:Book_id"`
	Borrow_book []Borrow_book `gorm:"foreignKey:Book_id"`

	Catagory_id *uint
	Catagory    Catagory `gorm:"references:id"`
}
type Borrow_status struct {
	gorm.Model
	Name string

	Borrow_book []Borrow_book `gorm:"foreignKey:Borrow_status_id"`
}
type Borrow_book struct {
	gorm.Model

	Start_time time.Time
	End_time   time.Time
	Totalbook  int

	Borrow_status_id int           //fk
	Borrow_status    Borrow_status `gorm:"references:id"`

	Member_id uint    //fk
	Member    *Member `gorm:"references:id"`

	Book_id uint  //fk
	Book    *Book `gorm:"references:id"`

	Point_history []Point_history `gorm:"foreignKey:Borrow_book_id"`
}
type Review struct {
	gorm.Model
	Rate    string
	Comment string

	Book_id uint  //fk
	Book    *Book `gorm:"references:id"`

	Member_id uint    //fk
	Member    *Member `gorm:"references:id"`
}
