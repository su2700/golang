package _case
import (
	"fmt"
	"os"
	"path"
)
func ReadWrite() {
   list := getFiles(sourceDir)
   for _, file := range list {
    byt, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	_, name:= path.Split(file)
	destName := destDir + "normal/" + name
	fmt.Println(destName)
	err = os.WriteFile(destName, byt, 0644)
	if err != nil {
		fmt.Println(err)
	}  
}
}
