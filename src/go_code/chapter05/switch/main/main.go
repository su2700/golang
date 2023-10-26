package main
import "fmt"
func main(){
     var key byte
	 fmt.Println("enter a char from a, b, c,d,e,f,g")
	 fmt.Scanf("%c", &key)
	 switch key{
	 	case 'a': fmt.Printf("1")
	 	case 'b': fmt.Printf("2")
		 case 'c': fmt.Println("3")
		 case 'd': fmt.Printf("4")
		 case 'e': fmt.Printf("5")
	 }

	 var num int =10
	 switch num {
	 case 10:
		fmt.Println("Ok1")
		fallthrough
	 case 20: // fallthrought will not check case after, exec directly 
		fmt.Println("Ok2")
	 case 30:
		fmt.Println("ok3")


	 }

}