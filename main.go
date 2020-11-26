package main

import (
	"github.com/gin-gonic/gin"
	"todo/todo"
)

func main() {
	setUpDB()
	setupRouter().Run(":8080")
}

func setupRouter() *gin.Engine {

	router := gin.Default()
	dbConn := getDb()

	router.GET("/todos", todo.Index(dbConn))
	router.GET("todos/:id/show", todo.Show(dbConn))
	router.POST("todos", todo.Create(dbConn))
	router.DELETE("todos/:id", todo.Delete(dbConn))
	router.PATCH("todos/:id/done", todo.ToggleDone(dbConn))

	return router
}
