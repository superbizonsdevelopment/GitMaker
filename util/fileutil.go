package util

import(
	"log"
	"os"
)

func MakeAppDir() {
	log.Println("system dir: ", os.Getenv("windir"))
	AppDir := "Users/Antoni/Desktop/GitMaker/"
	log.Println(AppDir)
	
	_, err = os.Stat(AppDir)
	
	if os.IsNotExist(err) {
		file, err := os.Create(AppDir)
		
		if err != nil {
			panic(err.Error())
		}
		defer file.Close()
	}
	
	log.Println("Created Maing app's dir!!!")
}