package router

import (
  "encoding/json"
  "fmt"
  "io"
  "net/http"
)

type Router struct {
  Handlers *List
}

type Handler struct {
  Function func(ResponseWriter, *Request)
  Type string
  Route string
}

type ResponseWriter struct {
  http.ResponseWriter
}

type Request struct {
  *http.Request
  Params map[string]string
  Body   map[string]string
}

func (w *ResponseWriter) JSON(v interface{}) {
  json, err := json.MarshalIndent(v, "", " ")
  if err != nil {
    panic(err)
  }

  io.WriteString(w, string(json))
}

func CreateRouter() Router {
  return Router{Handlers: CreateList()}
}

func (r *Router) NotFound(w http.ResponseWriter, url, method string) {
  message := fmt.Sprintf("Cannot %s %s", method, url)
  w.WriteHeader(http.StatusNotFound)
  response := map[string]string{
    "message": string(message),
    "status":  fmt.Sprintf("%v", http.StatusNotFound),
  }
  json, err := json.MarshalIndent(response, "", " ")
  if err != nil {
    panic(err)
  }
  io.WriteString(w, string(json))
}

/* sets a handler for a certain method */
func (r *Router) setHandler(method, url string, handler func(ResponseWriter, *Request)) {
  h := Handler{Type: method, Function: handler, Route: url}
  r.Handlers.Push(h)
}

/* sets another GET handler */
func (r *Router) GET(url string, handler func(ResponseWriter, *Request)) {
  r.setHandler("GET", url, handler)
}

/* sets another POST handler */
func (r *Router) POST(url string, handler func(ResponseWriter, *Request)) {
  r.setHandler("POST", url, handler)
}

/* sets another PUT handler */
func (r *Router) PUT(url string, handler func(ResponseWriter, *Request)) {
  r.setHandler("PUT", url, handler)
}

/* sets another DELETE handler */
func (r *Router) DELETE(url string, handler func(ResponseWriter, *Request)) {
  r.setHandler("DELETE", url, handler)
}

func (r *Router) USE(handler func(ResponseWriter, *Request)) {
  r.setHandler("MIDDLEWARE", "", handler)
}

/* a root handler */
func (r Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  method := req.Method

  var handlerIndex int = -1
  for i, handler := range r.Handlers.ToArray() {
    if Match(handler.Route, req.URL.Path) && handler.Type == method {
      handlerIndex = i
      break
    }
  }

  if handlerIndex == -1 {
    r.NotFound(res, req.URL.Path, method)
    return
  }

  handler := r.Handlers.Get(handlerIndex)

  params := GetParams(handler.Route, req.URL.Path)
  body := make(map[string]string)
  if method == "POST" {
    body = ParseBody(req)
  }

  request := Request{req, params, body}
  for i := 0; i < handlerIndex; i++ {
    if handler := r.Handlers.Get(i); handler.Type != "MIDDLEWARE" {
      continue
    } else {
      handler.Function(ResponseWriter{res}, &request)
    }
  }

  handler.Function(ResponseWriter{res}, &request)
}





















