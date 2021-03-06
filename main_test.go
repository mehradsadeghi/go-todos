package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"todo/todo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var router *gin.Engine

func setup() {
	setUpDB()
	router = setupRouter()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestIfIndexIsWorkingCorrectly(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var todos []todo.Todo
	getDb().First(&todos)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestIfTodoCanBeShown(t *testing.T) {
	getDb().Create(todo.New("meh", "rad"))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/todos/1/show", nil)
	assert.Nil(t, err)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var todos []todo.Todo
	getDb().First(&todos)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestIfItCanCreate(t *testing.T) {
	w := httptest.NewRecorder()

	postBody, error := json.Marshal(todo.New("meh", "rad"))
	assert.Nil(t, error)

	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(postBody))
	require.Nil(t, err)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestIfTodoCanBeDeleted(t *testing.T) {
	item := todo.New("meh", "rad")
	db.Save(item)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/todos/" + fmt.Sprint(item.Id), nil)
	require.Nil(t, err)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var model todo.Todo
	getDb().Delete(&model, item.Id)
}

func TestIfDoneCanBeToggled(t *testing.T) {
	item := todo.New("meh", "rad")
	db.Save(item)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", "/todos/" + fmt.Sprint(item.Id) + "/done", nil)
	require.Nil(t, err)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var model todo.Todo
	getDb().Delete(&model, item.Id)
}
