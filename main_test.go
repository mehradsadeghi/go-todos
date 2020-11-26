package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIfIndexIsWorkingCorrectly(t *testing.T) {

	setUpDB()

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var todos []Todo
	getDb().First(&todos)

	assert.Equal(t, len(todos), 0)
}

func TestIfItCanStore(t *testing.T) {

	setUpDB()

	router := setupRouter()

	w := httptest.NewRecorder()

	data := Todo{Title: "mehrad", Body: "sadeghi"}

	req, _ := http.NewRequest("POST", "/todos", data)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var todos []Todo
	getDb().First(&todos)

	assert.Equal(t, len(todos), 1)

	getDb().Delete(todos)
}

func (Todo) Read(p []byte) (n int, err error) {
	p, _ = json.Marshal(p)
	return len(p), nil
}