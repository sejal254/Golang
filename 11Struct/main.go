package main

import "fmt"

func main() {

	fmt.Println("Welocme to Struct")

	//in struct no inheritance in golang;
	//No concept of super or parent class

	//lets use struct in main 
	sejal :=User{"Sejal","sejal@gmail.com",true,21}
	fmt.Println(sejal)
	fmt.Printf("sejal details are: %+v\n",sejal)
	fmt.Printf("Name is %v and email is %v.",sejal.Name,sejal.Email)
}

//define a structure
type User struct{
	Name string
	Email string
	Status bool
	Age int
}