package rest

import (
  "encoding/json"
  "fmt"
  "io"
  "net/http"
)

type Router struct {
  Routes map[string]map[string]func(ResponseWriter, *Request)
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
  return Router{}
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
  if r.Routes == nil {
    r.Routes = make(map[string]map[string]func(ResponseWriter, *Request))
    r.Routes[method] = make(map[string]func(ResponseWriter, *Request))
  }

  if r.Routes[method] == nil {
    r.Routes[method] = make(map[string]func(ResponseWriter, *Request))
  }

  r.Routes[method][url] = handler
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

/* a root handler */
func (r Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  method := req.Method

  for url, handler := range r.Routes[method] {
    if Match(url, req.URL.Path) {
      params := GetParams(url, req.URL.Path)
      request := Request{req, params, make(map[string]string)}

      if method == "POST" {
        request.Body = ParseBody(req)
      }

      handler(ResponseWriter{res}, &request)

      return
    }
  }

  r.NotFound(res, req.URL.Path, method)
}
