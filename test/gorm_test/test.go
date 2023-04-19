package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   int `gorm:"primary_key"`
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test_db"))
	if err != nil {
		panic(err)
	}
	err = db.Table("users").Where("id=?", 4).Update("agee", "5").Error
	fmt.Println(err)
	//db.AutoMigrate(&User{})
	//userList := []User{
	//	{ID: 1, Name: "Name1", Age: 11},
	//	{ID: 2, Name: "Name2", Age: 12},
	//	{ID: 3, Name: "Name3", Age: 13},
	//}
	//for _, user := range userList {
	//	db.Create(&user)
	//}
}
