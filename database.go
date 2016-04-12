package main

import "github.com/jinzhu/gorm"

var databaseConnection *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "/tmp/techtalk.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(User{})
	user := User{
		FirstName: "Linh",
		LastName:  "Nguyen",
	}
	db.FirstOrCreate(&user)
	db.Debug()
	databaseConnection = db
}
