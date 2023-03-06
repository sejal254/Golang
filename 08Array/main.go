package main

import "fmt"

func main() {
	fmt.Println("welcome to array")

	//note:- very less use of array in go

//decalre an array
    var fruitList [4] string//complusary provide number 
	fruitList[0]="apple"
	fruitList[1]="orange"
	fruitList[3]="peach"
	fmt.Println("FruitList",fruitList)
	fmt.Println("FruitList",len(fruitList))

	var vegList =[3]string{"potato","tomato","carrot"}
	fmt.Println("Vegies List", vegList)



}