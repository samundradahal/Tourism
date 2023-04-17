package main

import (
	"tourism/db"
	model "tourism/model"
)

func init() {
	db.ConnectDb()
}

func main() {
	db.DB.AutoMigrate(&model.Tourism{})
}
