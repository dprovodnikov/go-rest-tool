package main

import (
  "restful/rest"
  "net/http"
  "io"
)

func foo(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "bar")
}

func main() {
  router := rest.Router{}

  router.GET("/foo", foo)
  router.GET("/bar", bar)

  http.ListenAndServe(":8080", router)
}