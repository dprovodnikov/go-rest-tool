package main

import (
  "restful/rest"
  "net/http"
  "io"
)

func foo(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "foo get")
}

func bar(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "foo post")
}


func main() {
  router := rest.Router{}

  router.GET("/foo", foo)
  router.POST("/foo", bar)

  http.ListenAndServe(":8080", router)
}