package main

import (
  "net/http"
  "restful/rest"
)

func handler(w rest.ResponseWriter, r *rest.Request) {
  w.JSON(map[string]string{
    "first_name": r.Params["f_name"],
    "last_name":  r.Params["l_name"],
  })
}

func main() {
  router := rest.Router{}

  router.POST("/user/:f_name/:l_name", handler)

  http.ListenAndServe(":8080", router)
}
