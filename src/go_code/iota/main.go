package main

import (
	"fmt"
//	"internal/itoa"
)

func main(){

	const(

		A1=1<<iota
		A2
		A3
		A4

	)
	fmt.Println(A1,A2,A3,A4)

}