package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	var todos []Todo
	getDb().Find(&todos)
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

	if result := getDb().First(&todo, id); result.Error != nil {
		c.JSON(http.StatusOK, "Item not found")
		return
	}

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

	getDb().Create(&Todo{Title: json.Title, Body: json.Body})

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func toggleUnDone(c *gin.Context) {

	var json struct {
		ID string `json:"id" url:"required"`
	}

	if c.Bind(&json) != nil {
		c.JSON(http.StatusOK, json.ID+" is not valid")
		return
	}

	id := c.Params.ByName("id")

	var todo Todo

	getDb().First(&todo, id)

	if result := getDb().First(&todo, id); result.Error != nil {
		c.JSON(http.StatusOK, "Item not found")
		return
	}

	getDb().Model(&todo).Update("Done", false)

	c.JSON(http.StatusOK, "OK")
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

	getDb().First(&todo, id)

	if result := getDb().First(&todo, id); result.Error != nil {
		c.JSON(http.StatusOK, "Item not found")
		return
	}

	getDb().Model(&todo).Update("Done", true)

	c.JSON(http.StatusOK, "OK")
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

	if result := getDb().Delete(&todo, id); result.Error != nil {
		c.JSON(http.StatusOK, "Item not found")
		return
	}
}

