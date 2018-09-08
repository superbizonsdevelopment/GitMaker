package main

import (
	"./constants"
	"./handler"
	"./util"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Gitmaker Code Compilation Server!!!")
	log.Println("Creating Directories!")
	util.CreateDirIfNotExist(constants.AppDir)
	util.CreateDirIfNotExist(constants.AppsDir)
	log.Println("All Directories created!")

	mux := http.NewServeMux()

	mux.HandleFunc("/gccs", handler.HandleMainFunction)

	http.ListenAndServe(":8080", mux)
}
