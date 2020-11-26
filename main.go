package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Todo struct {
	ID        uint         `json:"id" gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time    `json:created_at`
	UpdatedAt time.Time    `json:updated_at`
	Title     string       `json:"title"`
	Body      string       `json:"body"`
	Done      sql.NullBool `json:"done" gorm:"default:false"`
}

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

func setupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("todos/", index)
	router.GET("todos/:id", show)
	router.POST("todos", store)
	router.DELETE("todos/:id", delete)

	router.PATCH("todos/:id/done", toggleDone)

	return router
}

func toggleDone(c *gin.Context) {

	var json struct {
		ID string `json:"id" url:"required"`
	}

	if c.Bind(&json) != nil {
		c.JSON(http.StatusOK, json.ID+" is not valid")
		return
	}

	id := c.Params.ByName("id")

	var todo Todo

	db.First(&todo, id)

	db.Model(&todo).Update("Done", true)

	c.JSON(http.StatusOK, "OK")
}

func index(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func show(c *gin.Context) {

	var json struct {
		ID string `json:"id" url:"required"`
	}

	if c.Bind(&json) != nil {
		c.JSON(http.StatusOK, json.ID+" is not valid")
		return
	}

	id := c.Params.ByName("id")

	var todo Todo
	db.First(&todo, id)

	c.JSON(http.StatusOK, gin.H{"Title": todo.Title, "Body": todo.Body})
}

func store(c *gin.Context) {

	var json struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}

	if c.Bind(&json) != nil {
		c.JSON(http.StatusUnprocessableEntity, "Provided data is not valid")
		return
	}

	db.Create(&Todo{Title: json.Title, Body: json.Body})

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func delete(c *gin.Context) {

	var json struct {
		ID string `json:"id" url:"required"`
	}

	if c.Bind(&json) != nil {
		c.JSON(http.StatusOK, json.ID+" is not valid")
		return
	}

	id := c.Params.ByName("id")

	var todo Todo
	db.Delete(&todo, id)
}

func main() {
	setUpDB()
	setupRouter().Run(":8080")
}
