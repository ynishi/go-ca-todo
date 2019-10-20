// Package repository implements simple in memory repositories.
// It is for development or test, so in the internal pkg.
// Internal pkg can hide repository implementation from out of hole package,
// but it is need to public when inject repository at setup app.
// So repository implementation mainly set to pkg/todo/repository or pkg/repository directory.
package repository

import (
	"github.com/ynishi/go-ca-todo/pkg/todo"
)

// InMemoryTodoContent is a concrete implementation of TodoContent interface.
type InMemoryTodoContent struct {
	todos map[string]todo.Todo
}

// FetchAll returns all Todos from private variable.
func (tc InMemoryTodoContent) FetchAll() []todo.Todo {
	var ts []todo.Todo
	for _, value := range tc.todos {
		ts = append(ts, value)
	}
	return ts
}

// Add creates new Todo and save to private variable.
func (tc *InMemoryTodoContent) Add(t *todo.Todo) {
	if tc.todos == nil {
		tc.todos = make(map[string]todo.Todo)
	}
	tc.todos[t.Title] = *t
}

// InMemoryTag is a concrete implementation of Tag interface.
type InMemoryTag struct {
	tags map[string]string
}

// GetByTitle returns tag matches title.
func (t InMemoryTag) GetByTitle(title string) string {
	return t.tags[title]
}

// Add creates new Tag and save to private variable.
func (t *InMemoryTag) Add(tag, title string) {
	if t.tags == nil {
		t.tags = make(map[string]string)
	}
	t.tags[title] = tag
}
