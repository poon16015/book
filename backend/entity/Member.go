package entity

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	FirstName string
	LastName string
	Tel int
	Email string
	Password string
	point float32

	

}