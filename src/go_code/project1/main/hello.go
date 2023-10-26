package main // hello.go belong package main, all go file must belong a package
import (
	"fmt"
	"time"

)
func main(){
	go fmt.Println("Hello World1") //This line of code starts with the keyword go, which is used to launch a new goroutine in Go. A goroutine is a lightweight thread of execution.
	go fmt.Println("Hello World2") //goroutine , concurrency, no order
	go fmt.Println("Hello World3") 
	go fmt.Println("Hello World4") 
	var num=10
	fmt.Println(num)
	time.Sleep(time.Second)


}


// hello.go belong to the package main
// func main, main is name of func, 
// in go, all file must belong to a package
//in go, it run line one by one, and semicolon is unnecessary 
//escape char
// \t  table, \r return



