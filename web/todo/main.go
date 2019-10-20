// Package main in web implements http server application,
// the most outer layer of architecture same as cmd.
// This contains server init process and register handler
// run parser, use case, presenter flow.
package main

import (
	"github.com/ynishi/go-ca-todo/internal/pkg/repository"
	"github.com/ynishi/go-ca-todo/pkg/todo/adapter"
	"github.com/ynishi/go-ca-todo/pkg/todo/interacter"
	wAdapter "github.com/ynishi/go-ca-todo/web/todo/adapter"
	"log"
	"net/http"
)

var (
	todoInteracter *interacter.TodoInteracter
)

func init() {
	todoInteracter = interacter.NewTodoIntercter(&repository.InMemoryTodoContent{}, &repository.InMemoryTag{})
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	res := todoInteracter.List()
	adapter.RunPresent(adapter.EchoPresenterWith(w), *res)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	req := adapter.RunParse(wAdapter.QueryParse(r), nil)
	res := todoInteracter.Add(req)
	adapter.RunPresent(adapter.EchoPresenterWith(w), *res)
}

func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/add", addHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
