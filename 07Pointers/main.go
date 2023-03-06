package main

import "fmt"

func main() {
	fmt.Println("welcome to Pointers")

//pointer pointing nothing	
	// var ptr *int
	// fmt.Println("value of pointers is:",ptr)

//pointers point something
	mynumber:=23
	var ptr = &mynumber
	fmt.Println("value of pointer is",ptr)
	fmt.Println("value of actual pointer is",*ptr)

	*ptr=*ptr*2
	fmt.Println("New valur is:",mynumber)
	fmt.Println("New valur is:",*ptr)


}