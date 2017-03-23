package main

import (
  "restful/rest"
  "net/http"
  "io"
  "fmt"
)

func handler(w rest.ResponseWriter, r *rest.Request) {
  output := fmt.Sprintf("id is:(%s)\nmode is:(%s)\n", r.Params["id"], r.Params["mode"])
  io.WriteString(w, output)
}


func main() {
  router := rest.Router{}

  router.GET("/user/:id/dashboard/:mode", handler)

  http.ListenAndServe(":8080", router)
}