package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")
	greeter()
	result :=adder(3,5)
	fmt.Println("Result:",result)
}

func greeter(){
	fmt.Println("Namaste")
}

func adder(valOne int, valtwo int) int {//return type
	return valOne+valtwo

} 

func proAdder(values ...int) (int,string){
	total :=0

	for _ , val:=range values{
		total+=val

	}
	return total,"Hi pro result Function"
}