package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	setUpDB()
	setupRouter().Run(":8080")
}

func setupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("todos/", index)
	router.GET("todos/:id", show)
	router.POST("todos", store)
	router.DELETE("todos/:id", delete)

	router.PATCH("todos/:id/done", toggleDone)
	router.PATCH("todos/:id/undone", toggleUnDone)

	return router
}
