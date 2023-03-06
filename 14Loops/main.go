package main

import "fmt"

func main() {
	fmt.Println("welcome to Loops")

	//create a slcie of days
	days := []string{"Sunday","Monday","Tusday","Wednesday","Thursday","Friday","Saturday"}
	fmt.Println(days)

	// for d:=0;d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	for i:=range days{
		fmt.Println(days[i])
	}

	//in loops we also use break, continue,goto
}