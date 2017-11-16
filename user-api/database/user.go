package database

import (
	"github.com/jinzhu/gorm"
)

// User type
type User struct {
	gorm.Model
	Name  string
}

func initUser() {
	// create table if it does not exist
	if !DB.HasTable(&User{}) {
		DB.CreateTable(&User{})
		testUser := User{Name: "Test"}
		DB.Create(&testUser)
	}
}