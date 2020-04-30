// callback router
package nap

import (
  "net/http"
  "fmt"
)

type RouterFunc func(resp *http.Response , content interface{}) error // return a function (toggle if necessary)

type CBRouter struct {
  Routers map[int]RouterFunc // int = statusCode <> function execution
  DefaultRouter RouterFunc
}

// NewRouter returns a new router
func NewRouter() *CBRouter {
	return &CBRouter{
		Routers: make(map[int]RouterFunc),
		DefaultRouter: func(resp *http.Response, _ interface{}) error {
			return fmt.Errorf("From: %s received unknown status: %d",
				resp.Request.URL.String(), resp.StatusCode)
		},
	}
}


func (r *CBRouter) RegisterFunc(status int, fn RouterFunc) {
	r.Routers[status] = fn
}

func (r *CBRouter) CallFunc(resp *http.Response, content interface{}) error {
	fn, ok := r.Routers[resp.StatusCode]
	if !ok {
		fn = r.DefaultRouter
	}
	if err := fn(resp, content); err != nil {
		return err
	}
	return nil
}
