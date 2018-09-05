package main

import(
  "log"
  "./util"
  "net/http"
  "./constants"
)

func main() {
  log.Println("Starting Gitmaker Code Compilation Server!!!")
  log.Println("Creating Directories!")
  util.CreateDirIfNotExist(constants.AppDir)
  util.CreateDirIfNotExist(constants.AppsDir)
  log.Println("All Directories created!")

  mux := http.NewServeMux()

  mux.HandleFunc("/gccs", func(w http.ResponseWriter, r *http.Request){

    switch r.Method {
    case "GET":
      break
    case "POST":

    }
  })

  http.ListenAndServe(":8080", mux)
}
