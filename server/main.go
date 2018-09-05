package main

import(
  "log"
  "net/http"
  "./util"
)

const(
  AppDir = "/Applications/GitMaker-Server"
  AppsDir = "/Applications/GitMaker-Server/apps"
)

func main() {
  log.Println("Starting Gitmaker Code Compilation Server!!!")
  log.Println("Creating Directories!")
  util.CreateDirIfNotExist(AppDir)
  util.CreateDirIfNotExist(AppsDir)
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
