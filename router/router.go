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

type Data map[string]interface{}
type HandlerFunc func(ResponseWriter, *Request)
type Handler struct {
  Function HandlerFunc
  Type string
  Route string
}
type MiddlewareFunc func(ResponseWriter, *Request) int
type Middleware struct {
  Function MiddlewareFunc
  Type string
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
func (r *Router) setHandler(method, url string, handler HandlerFunc) {
  h := Handler{Type: method, Function: handler, Route: url}
  r.Handlers.Push(h)
}

/* sets another GET handler */
func (r *Router) GET(url string, handler HandlerFunc) {
  r.setHandler("GET", url, handler)
}

/* sets another POST handler */
func (r *Router) POST(url string, handler HandlerFunc) {
  r.setHandler("POST", url, handler)
}

/* sets another PUT handler */
func (r *Router) PUT(url string, handler HandlerFunc) {
  r.setHandler("PUT", url, handler)
}

/* sets another DELETE handler */
func (r *Router) DELETE(url string, handler HandlerFunc) {
  r.setHandler("DELETE", url, handler)
}

func (r *Router) USE(handler MiddlewareFunc) {
  m := Middleware{Type: "MIDDLEWARE", Function: handler}
  r.Handlers.Push(m)
}

/* a root handler */
func (r Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  method := req.Method

  var handlerIndex int = -1
  SearchingForIndex:
    for i, handler := range r.Handlers.ToArray() {
      switch handler := handler.(type) {
      default:
        continue SearchingForIndex
      case Handler:
        if Match(handler.Route, req.URL.Path) && handler.Type == method {
          handlerIndex = i
          break SearchingForIndex
        }
      }
    }

  if handlerIndex == -1 {
    r.NotFound(res, req.URL.Path, method)
    return
  }

  handler := r.Handlers.Get(handlerIndex).(Handler)

  params := GetParams(handler.Route, req.URL.Path)
  body := make(map[string]string)
  if method == "POST" {
    body = ParseBody(req)
  }

  request := Request{req, params, body}
  SearchingForMiddlewares:
    for i := 0; i < handlerIndex; i++ {
      switch handler := r.Handlers.Get(i).(type) {
        default:
          continue SearchingForMiddlewares
        case Middleware:
          status := handler.Function(ResponseWriter{res}, &request)
          if status >= 300 {
            // Abort with the status code
            return
          }
      }
    }

  handler.Function(ResponseWriter{res}, &request)
}





















