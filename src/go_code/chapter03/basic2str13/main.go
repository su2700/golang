package main
import (
	"fmt"
	_ "unsafe"
)

func main()  {

	var n1 int =99
	var n2 float32=23.75
	// var b bool = true
	// var myChar byte = 'h'
	 var str string
	str = fmt.Sprintf("%d", n1)
	fmt.Printf("str type %T str=%v\n", str, str)

    

	
}