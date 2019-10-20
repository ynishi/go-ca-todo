package todo

import (
	"testing"
	"time"
)

var (
	cur = time.Now()
	due = cur.Add(time.Second * 10)
)

func TestNew(t *testing.T) {
	todo := NewTodo("t", "c", due, cur)
	expected := &Todo{
		"t",
		"c",
		due,
		cur,
		false,
	}
	if *todo != *expected {
		t.Errorf("Not matched: \n%v\n%v\n", todo, expected)
	}
}

func TestTodo_Done(t *testing.T) {
	todo := NewTodo("t", "c", due, cur)
	todo.Done()
	expected := &Todo{
		"t",
		"c",
		due,
		cur,
		true,
	}
	if *todo != *expected {
		t.Errorf("Not matched: \n%v\n%v\n", todo, expected)
	}
}
