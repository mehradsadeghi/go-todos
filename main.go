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
	router.GET("/todos", todo.Index(getDb()))
	router.GET("todos/:id/show", todo.Show(getDb()))
	router.POST("todos", todo.Create(getDb()))
	router.DELETE("todos/:id", todo.Delete(getDb()))
	router.PATCH("todos/:id/done", todo.ToggleDone(getDb()))

	return router
}
