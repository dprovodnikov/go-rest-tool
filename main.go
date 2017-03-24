package main

import (
  "fmt"
  "io"
  "net/http"
  "restful/rest"
)

func handler(w rest.ResponseWriter, r *rest.Request) {
  io.WriteString(w, "Hello world")
}

func main() {
  router := rest.Router{}

  router.POST("/user/:id", handler)

  http.ListenAndServe(":8080", router)
}
