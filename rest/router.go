package rest

import (
  "net/http"
)

type Router struct {
  Routes map[string]map[string]func(http.ResponseWriter, *http.Request)
}

/* sets a handler for a certain method */
func (r *Router) setHandler(method, url string, handler func(http.ResponseWriter, *http.Request)) {
  if r.Routes == nil {
    r.Routes = make(map[string]map[string]func(http.ResponseWriter, *http.Request))
    r.Routes[method] = make(map[string]func(http.ResponseWriter, *http.Request))
  }

  if r.Routes[method] == nil {
    r.Routes[method] = make(map[string]func(http.ResponseWriter, *http.Request))
  }

  r.Routes[method][url] = handler
}

/* sets another GET handler */
func (r *Router) GET(url string, handler func(http.ResponseWriter, *http.Request)) {
  r.setHandler("GET", url, handler)
}

/* sets another POST handler */
func (r *Router) POST(url string, handler func(http.ResponseWriter, *http.Request)) {
  r.setHandler("POST", url, handler)
}


/* sets another PUT handler */
func (r Router) PUT(url string, handler func(http.ResponseWriter, *http.Request)) {
  r.setHandler("PUT", url, handler)
}

/* sets another DELETE handler */
func (r Router) DELETE(url string, handler func(http.ResponseWriter, *http.Request)) {
  r.setHandler("DELETE", url, handler)
}

/* a root handler */
func (r Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  method := req.Method

  for url, handler := range r.Routes[method] {
    if req.URL.Path == url {
      handler(res, req)
    }
  }
}