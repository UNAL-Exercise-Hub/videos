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
	dsn := "unworkout:unworkout@tcp(127.0.0.1:3306)/videos_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Database connected")
	}
	if !db.Migrator().HasTable("videos") {
		createDB()
	}
}
