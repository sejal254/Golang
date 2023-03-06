package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")
	var fruitList =[]string{"apple", "mango", "banana"}
	fmt.Printf("Type of fruit: %T \n",fruitList)

	fruitList=append(fruitList, "Peach","watermelon")
	fmt.Println(fruitList)


	fruitList=append(fruitList[1:3])
	fmt.Println(fruitList)


	//Advantage for using append method
	highscore:=make([]int,4)

	highscore[0]=234
	highscore[1]=342
	highscore[2]=453
	highscore[3]=511

	highscore=append(highscore, 555,666,111)
	//append reloacte the memory and will all of the value accomadated

	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)


	//how to remove a vale from slices based on index

	var courses =[]string{"react","java","swift","python","ruby"}
	fmt.Println(courses)

	//we can also use append to remove the values
	var index int=2
	courses=append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
	




}