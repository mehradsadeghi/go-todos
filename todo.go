package main

import (
	"database/sql"
	"time"
)

type Todo struct {
	Id        uint         `json:"id" gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Title     string       `json:"title"`
	Body      string       `json:"body"`
	Done      sql.NullBool `json:"done" gorm:"default:false"`
}
