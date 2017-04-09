package main

import (
  "rest/router"
  "net/http"
)

func handler(w router.ResponseWriter, req *router.Request) {
  w.JSON(router.Data{
    "some": req.Params["some"],
    "username": req.Body["username"],
    "password": req.Body["password"],
  })
}

func middleware(w router.ResponseWriter, req *router.Request) {
  w.JSON(router.Data{"message": "middleware :)"})
}

func main() {
  r := router.CreateRouter()

  r.USE(middleware)
  r.USE(middleware)
  r.USE(middleware)
  r.POST("/:some/", handler)

  http.ListenAndServe(":8080", r)
}