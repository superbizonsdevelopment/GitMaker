package util

import (
	"log"
	"os"
)

func CreateDirIfNotExist(dir string) {
	src, err := os.Stat(dir)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dir, 0755)

		if errDir != nil {
			panic(nil)
		}

		log.Printf("Created %s package!", dir)
	}

	src, err = os.Stat(dir)

	if src.Mode().IsRegular() {
		log.Printf(dir, " already exist as a file!")
	}
}
