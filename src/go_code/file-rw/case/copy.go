package _case
import (
	"path"
	"os"
	"io"
	"log"
)

func CopyDirToDir()  {

   list := getFiles(sourceDir)
   for _, f := range list {
       _, name := path.Split(f) //get file name
	   destFileName := destDir + "copy/"+ name //copy file
	   COPYFile(f, destFileName)
   }

}

func COPYFile(srcName, destName string) (w int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
		}
		defer src.Close()
		dst, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
		log.Fatal(err)
		}
		defer dst.Close()
		return io.Copy(dst, src)
}