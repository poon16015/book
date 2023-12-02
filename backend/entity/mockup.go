package entity

func MockupAdmin(){
	admin1 := Admin{
		Email : "admin@gmail.com",
		Password: "admin1234",
	}
	db.Model(&Admin{}).Create(&admin1)
}
func MockupRoom_type(){
	Rtype1 := Room_type{
		Name : "small",
		Capacity : 5 ,
	}
	db.Model(&Room_type{}).Create(&Rtype1)
	Rtype2 := Room_type{
		Name : "medium",
		Capacity : 10 ,
	}
	db.Model(&Room_type{}).Create(&Rtype2)
	Rtype3 := Room_type{
		Name : "large",
		Capacity : 20 ,
	}
	db.Model(&Room_type{}).Create(&Rtype3)
}