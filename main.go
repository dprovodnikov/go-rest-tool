package main

import (
  rest "rest/router"
  "net/http"
)

func handler(w rest.ResponseWriter, req *rest.Request) {
  w.Abort(403, http.StatusText(403))

  w.JSON(rest.Data{
    "some": req.Params["some"],
    "username": req.RequestBody["username"],
    "password": req.RequestBody["password"],
  })
}

func middleware(w rest.ResponseWriter, req *rest.Request) int {
  if (req.RequestBody["username"] == "wireden") {
    w.JSON(rest.Data{"secret-key": "1111"})
    return http.StatusOK
  } else {
    return http.StatusInternalServerError
  }
}

func main() {
  r := rest.CreateRouter()

  r.USE(middleware)
  r.POST("/:some", handler)

  http.ListenAndServe(":8080", r)
}