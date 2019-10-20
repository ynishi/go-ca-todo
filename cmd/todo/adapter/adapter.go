// Package adapter in cmd implements interface adapter
// for cmd specified implementation.
package adapter

import (
	"fmt"
	"github.com/ynishi/go-ca-todo/pkg/todo/adapter"
	"github.com/ynishi/go-ca-todo/pkg/todo/interacter"
	"os"
)

// SubParse create Parser insert sub command to request.
func SubParse(sub string) adapter.Parser {
	return func(c fmt.Scanner) (interacter.Req, fmt.Scanner) {
		req := interacter.Req{
			Sub: sub,
		}
		return req, c
	}
}

// EnvSubParser create Parser do EnvParse and SubParse.
func EnvSubParser(sub string) adapter.Parser {
	return adapter.All(adapter.EnvParse(), SubParse(sub))
}

// EchoPresenter create Parser that print res to stdout
func EchoPresenter() adapter.Presenter {
	return adapter.EchoPresenterWith(os.Stdout)
}
