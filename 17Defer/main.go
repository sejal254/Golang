package main

import "fmt"

func main() {
	// fmt.Println("Hello")
	// defer fmt.Println("world")

	// defer fmt.Println("world")
	// fmt.Println("Hello")
	//when we use deffer that function wiil execute at the last


	defer fmt.Println("world")
	defer fmt.Println("Hello")
	defer fmt.Println("My")
	//many defer will follow LIFO rule


}

func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Print(i)
	}
}