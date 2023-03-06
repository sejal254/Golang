package main

import "fmt"

func main() {
	var username string = "sejal"
	fmt.Println(username)
	//T is place holder in Go
	fmt.Printf("type of variable : %T \n", username)
	var smallint int16 = 250

	//boolean
	var isLoggedIn  bool =false
	fmt.Println(isLoggedIn)
	fmt.Printf("type of variable :%T \n",isLoggedIn)


	fmt.Println(smallint)
	fmt.Printf("type of variable : %T \n", smallint)
	 
	var smallfloat float32 = 250.32
	fmt.Println(smallfloat)
	fmt.Printf("type of variable : %T")
}
