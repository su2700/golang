package main
import (
	"fmt"
)

var (
   n7=100
   n8="tom2"

)
// var's vaule can be change in the same zone and same type
func main()  {
	var n1, n2 int
	var n3, n4 = 100, "tom"
	n5:= "jery"
	fmt.Println(n1, n2, n3, n4, n5)
}