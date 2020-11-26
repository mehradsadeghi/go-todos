package todo

import (
	"time"
)

type Todo struct {
	Id        uint      `json:"Id" gorm:"primaryKey;autoIncrement:true"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Done      bool      `json:"done" gorm:"default:false"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func New(title, body string) *Todo {
	return &Todo{
		Title: title,
		Body:  body,
	}
}
