package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

const DbFile = "todos.db"

func setUpDB() {

	var error error

	db, error = gorm.Open(sqlite.Open(DbFile), &gorm.Config{})

	if error != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})
}

func getDb() *gorm.DB {
	return db
}