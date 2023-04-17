package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var er error
	dsn := "root:root@tcp(127.0.0.1:3306)/details?parseTime=true"
	db, er := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if er != nil {
		panic("Error connecting Database")
	}
	DB = db
	fmt.Println("DataBase Connected")

}
