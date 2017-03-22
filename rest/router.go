package rest

import (
  "net/http"
)

type Router struct {
  Routes map[string]map[string]func(http.ResponseWriter, *http.Request)
}

func (r *Router) GET(url string, handler func(http.ResponseWriter, *http.Request)) {
  if r.Routes == nil {
    r.Routes = make(map[string]map[string]func(http.ResponseWriter, *http.Request))
    r.Routes["GET"] = make(map[string]func(http.ResponseWriter, *http.Request))
  }

  if r.Routes["GET"] == nil {
    r.Routes["GET"] = make(map[string]func(http.ResponseWriter, *http.Request))
  }

  r.Routes["GET"][url] = handler
}

func (r *Router) POST(url string, handler func(http.ResponseWriter, *http.Request)) {
  if r.Routes == nil {
    r.Routes = make(map[string]map[string]func(http.ResponseWriter, *http.Request))
    r.Routes["POST"] = make(map[string]func(http.ResponseWriter, *http.Request))
  }

  if r.Routes["POST"] == nil {
    r.Routes["POST"] = make(map[string]func(http.ResponseWriter, *http.Request))
  }

  r.Routes["POST"][url] = handler
}

func (r Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  method := req.Method

  for url, handler := range r.Routes[method] {
    if req.URL.Path == url {
      handler(res, req)
    }
  }
}