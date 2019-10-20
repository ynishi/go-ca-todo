// Package adapter in web implements interface adapter
// for specified to web.
package adapter

import (
	"fmt"
	"github.com/ynishi/go-ca-todo/pkg/todo/adapter"
	"github.com/ynishi/go-ca-todo/pkg/todo/interacter"
	"net/http"
	"time"
)

// QueryParse create Parser convert url query parameter to request.
func QueryParse(r *http.Request) adapter.Parser {
	return func(c fmt.Scanner) (interacter.Req, fmt.Scanner) {
		q := r.URL.Query()
		fmt.Println(q)
		req := interacter.Req{}
		if q == nil {
			return req, c
		}
		if v, ok := q["title"]; ok {
			req.Title = v[0]
		}
		if v, ok := q["tag"]; ok {
			req.Tag = v[0]
		}
		if v, ok := q["due"]; ok {
			due, _ := time.Parse(time.RFC3339, v[0])
			req.Due = due
		}
		req.Sub = r.URL.Path
		return req, c
	}
}
