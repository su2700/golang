package _case
import (
	"log"
	"os"
	"strings"

)
const sourceDir = "source-file/"
const destDir = "dest-file/"


func getFiles(dir string) []string {
	fs, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)	}
	list := make([]string, 0)	
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		//fullName := strings.Trim(dir, cutset:"/" ) + "/" + f.Name()
		fullName := strings.Trim(dir, "/") + "/" + f.Name()

		list = append(list, fullName)
	}
	return list
}