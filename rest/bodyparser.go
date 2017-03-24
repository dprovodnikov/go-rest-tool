package rest

import (
  "net/http"
)

func ParseBody(r *http.Request) map[string]string {
  err := r.ParseForm()
  if err != nil {
    panic(err)
  }

  body := make(map[string]string)

  for key, value := range r.PostForm {
    body[key] = value[0]
  }

  return body
}
