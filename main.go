package main

import (
  "net/http"
  "restful/rest"
)

func handler(w rest.ResponseWriter, r *rest.Request) {
  w.JSON(map[string]string{
    "key":  "value",
    "key1": "value1",
  })
}

func main() {
  router := rest.Router{}

  router.POST("/user/:name", handler)

  http.ListenAndServe(":8080", router)
}
