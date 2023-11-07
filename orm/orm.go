package orm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Connect the orm with the database
func Connection() {
	dsn := "unworkout:unworkout@tcp(host.docker.internal:3307)/videos_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database\nPlease follow instructions in https://github.com/UNWorkout/Videos")
	} else {
		fmt.Println("Database connected")
	}
	if !db.Migrator().HasTable("videos") {
		createDB()
	}
}

// Get db from diferents packages
func GetDB() *gorm.DB {
	return db
}
