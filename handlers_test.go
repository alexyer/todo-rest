package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	Index(w, req)

	if w.Code != 200 || !strings.Contains(w.Body.String(), "Welcome!") {
		t.Errorf("Expected code '200' and body 'Welcome!', got code '%v', body '%v'", w.Code, w.Body)
	}
}

func TestTodoIndex(t *testing.T) {
	todos = todos[:0]
	todos = append(todos, Todo{Name: "Test Todo", Completed: true})

	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	TodoIndex(w, req)

	assert.Len(t, todos, 1)
	assert.Equal(t, todos[0].Name, "Test Todo")
	assert.True(t, todos[0].Completed)

	todos = append(todos, Todo{Name: "Another Todo", Completed: false})

	req, err = http.NewRequest("GET", "/todos", nil)
	if err != nil {
		log.Fatal(err)
	}

	w = httptest.NewRecorder()
	TodoIndex(w, req)

	assert.Len(t, todos, 2)
	assert.Equal(t, todos[1].Name, "Another Todo")
	assert.False(t, todos[1].Completed)
}

// FIXME(alexyer): Fix test later
func TestTodoShow(t *testing.T) {
	testCases := []int{1, 2}

	for _, testCase := range testCases {
		req, err := http.NewRequest("GET", "/todos/"+strconv.Itoa(testCase), nil)
		if err != nil {
			log.Fatal(err)
		}

		w := httptest.NewRecorder()
		TodoShow(w, req)

		assert.True(t, strings.Contains(w.Body.String(), "Todo show:"))
	}
}

func TestTodoCreate(t *testing.T) {
	todos = todos[:0]

	req, err := http.NewRequest("POST", "/todos", bytes.NewBufferString(`{"name": "Cool new note", "completed": false}`))
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	TodoCreate(w, req)

	assert.Len(t, todos, 1)
	assert.Equal(t, todos[0].Name, "Cool new note")
	assert.False(t, todos[0].Completed)
}
