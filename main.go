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

func a(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "a put")
}

func b(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "b delete")
}


func main() {
  router := rest.Router{}

  router.GET("/foo", foo)
  router.POST("/foo", bar)
  router.PUT("/foo", a)
  router.DELETE("/foo", b)

  http.ListenAndServe(":8080", router)
}