package main

import (
  "restful/rest"
  "net/http"
  "io"
)

func foo(w rest.ResponseWriter, r *rest.Request) {
  io.WriteString(w, "foo get")
}

func main() {
  router := rest.Router{}

  router.GET("/foo", foo)

  http.ListenAndServe(":8080", router)
}