// Package adapter implements part of interface adapter layer.
// This contains common interface of adapter and default implementation.
// It might be better to put this to client side(call use case like cli, web and so on.),
// but I think it is convenient to be provided unified interface.
package adapter

import (
	"fmt"
	interacter "github.com/ynishi/go-ca-todo/pkg/todo/interacter"
	"io"
	"os"
	"time"
)

// Parser(input to request boundary object, like controller in MVC)
// is implemented with higher-order function for flexibility.
type Parser func(fmt.Scanner) (interacter.Req, fmt.Scanner)

// All creates Parser that do all parsers and merge result.
func All(parsers ...interface{}) Parser {
	return func(c fmt.Scanner) (interacter.Req, fmt.Scanner) {
		acc := interacter.Req{}
		emptyDue := acc.Due
		for _, parser := range parsers {
			req, _ := doParse(parser, c)
			acc.Title += req.Title
			acc.Tag += req.Tag
			acc.Sub += req.Sub
			if !req.Due.Equal(emptyDue) {
				acc.Due = req.Due
			}
		}
		return acc, c
	}
}

// RunParse runs parser and returns parsed Req.
func RunParse(pi interface{}, c fmt.Scanner) interacter.Req {
	req, _ := doParse(pi, c)
	return req
}

func doParse(pi interface{}, c fmt.Scanner) (interacter.Req, fmt.Scanner) {
	switch p := pi.(type) {
	case Parser:
		return p(c)
	case *Parser:
		return (*p)(c)
	default:
		panic(fmt.Errorf("invalid type `%T`", pi))
	}
}

// Presenter(response boundary object to any view action like convert to view model
// or raise side effect).
// It is implemented with higher-order function for flexibility.
type Presenter func(interacter.Res) interacter.Res

// RunPresent runs presenter.
func RunPresent(presenter Presenter, r interacter.Res) {
	presenter(r)
}

// EnvParse return Parser that create req from Environment variable.
func EnvParse() Parser {
	return func(c fmt.Scanner) (interacter.Req, fmt.Scanner) {
		newTitle := os.Getenv("TITLE")
		newTag := os.Getenv("TAG")
		newDue := os.Getenv("DUE")
		due, err := time.Parse(time.RFC3339, newDue)
		if err != nil {
			panic(err)
		}
		req := interacter.Req{
			Title: newTitle,
			Tag:   newTag,
			Due:   due,
		}
		return req, c
	}
}

// EchoPresenterWith create Presenter that write res to w.
func EchoPresenterWith(w io.Writer) Presenter {
	return func(r interacter.Res) interacter.Res {
		buf := []byte(fmt.Sprintln(r))
		w.Write(buf)
		return r
	}
}
