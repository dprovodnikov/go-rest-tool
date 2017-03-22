package rest

import (
  "net/http"
)

type Router struct {
  Routes map[string]func(http.ResponseWriter, *http.Request)
}

func (r *Router) GET(url string, handler func(http.ResponseWriter, *http.Request)) {
  if r.Routes == nil {
    r.Routes = make(map[string]func(http.ResponseWriter, *http.Request))
  }

  r.Routes[url] = handler
}

func (r Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  handler := r.Routes[req.URL.Path]

  handler(res, req)
}