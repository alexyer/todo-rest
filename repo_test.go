package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoFindTodo(t *testing.T) {
	todos = todos[:0]
	todos = append(todos, Todo{Id: 42, Name: "Test Todo", Completed: true})

	todo := RepoFindTodo(42)

	assert.NotNil(t, todo)
	assert.Equal(t, todo.Name, "Test Todo")
	assert.True(t, todo.Completed)

	todo = RepoFindTodo(73)
	assert.NotNil(t, todo)
	assert.Equal(t, todo.Name, "")
	assert.False(t, todo.Completed)
}

func TestRepoDestroyTodo(t *testing.T) {
	todos = todos[:0]
	todos = append(todos, Todo{Id: 42, Name: "Test Todo", Completed: true})

	err := RepoDestroyTodo(42)
	assert.NoError(t, err)

	err = RepoDestroyTodo(42)
	assert.Error(t, err)

	err = RepoDestroyTodo(73)
	assert.Error(t, err)
}
