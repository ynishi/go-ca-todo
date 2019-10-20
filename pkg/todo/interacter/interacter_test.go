package interacter

import (
	"github.com/ynishi/go-ca-todo/pkg/todo"
	"testing"
	"time"
)

type TodoContentMock struct {
}

func (tc TodoContentMock) FetchAll() []todo.Todo {
	c := time.Now()
	d := c.Add(10 * time.Second)
	ret := todo.NewTodo("t1", "", d, c)
	return []todo.Todo{
		*ret,
	}
}

type TagMock struct {
}

func (t TagMock) GetByTitle(title string) string {
	return "t2"
}

func (tc TodoContentMock) Add(*todo.Todo) {
}

func (tc TagMock) Add(string, string) {

}

func TestTodoInteracter_List(t *testing.T) {
	ti := NewTodoIntercter(&TodoContentMock{}, &TagMock{})
	res := ti.List()
	if res.Todos[0].Title != "t1" {
		t.Error(res)
	}
	if res.Todos[0].Tag != "t2" {
		t.Error(res)
	}
}
