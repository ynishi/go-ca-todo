// Package main in cmd implements application executable,
// the most outer layer of architecture.
// This contains application init process and registry
// that implemented just switch selector.
package main

import (
	"flag"
	cAdapter "github.com/ynishi/go-ca-todo/cmd/todo/adapter"
	"github.com/ynishi/go-ca-todo/internal/pkg/repository"
	"github.com/ynishi/go-ca-todo/pkg/todo/adapter"
	interacter "github.com/ynishi/go-ca-todo/pkg/todo/interacter"
)

var (
	todoInteracter *interacter.TodoInteracter
	sub            string

	parser    adapter.Parser
	presenter adapter.Presenter
)

func init() {

	// init use case interacter
	todoInteracter = interacter.NewTodoIntercter(&repository.InMemoryTodoContent{}, &repository.InMemoryTag{})

	// parse input command(determine one use case).
	flag.StringVar(&sub, "sub", "usage", "sub cmd to run usecase")
	flag.Parse()

	// init interface adapters
	parser = cAdapter.EnvSubParser(sub)
	presenter = cAdapter.EchoPresenter()
}

func main() {

	// temporary registry.
	switch sub {
	case "usage":
		flag.Usage()
	case "list":
		res := todoInteracter.List()
		adapter.RunPresent(presenter, *res)
	case "add":
		req := adapter.RunParse(parser, nil)
		res := todoInteracter.Add(req)
		adapter.RunPresent(presenter, *res)
	}

}
