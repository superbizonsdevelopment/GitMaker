package handler

import (
  "net/http"
  "fmt"
)

func HandleMainFunction(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodGet {
    return
  }

  if r.Method == http.MethodPost {
    fmt.Fprintf(w,"Name: %v, SecondName: %v\n",r.PostFormValue("name"), r.PostFormValue("sname"))
  }
}
