// Package todo implements todo domain, contains Entity on top level of package.
// Todo Entity represents its business logic.
package todo

import "time"

//Todo is an Entity struct.
type Todo struct {
	Title  string
	Tag    string
	Due    time.Time
	Cur    time.Time
	IsDone bool
}

//NewTodo construct Todo with simple validation.
func NewTodo(title, tag string, due time.Time, cur time.Time) *Todo {
	if cur.After(due) {
		return nil
	}
	return &Todo{
		title,
		tag,
		due,
		cur,
		false,
	}
}

//Done update Todo ojbect to done state.
func (t *Todo) Done() {
	t.IsDone = true
}
