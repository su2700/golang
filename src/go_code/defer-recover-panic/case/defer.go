package _case

import "fmt"

// defer is FILO
func DeferCase1(){
  fmt.Println("start DeferCase1")
  defer func ()  {
	fmt.Println("Anonymous functions1 ")
  }()
  defer f1()
  defer func () {
	fmt.Println("Anonymous functions2 ")
  } ()

  fmt.Println("end")
}

func f1()  {
	fmt.Println("func	f1 ")
  
	
}