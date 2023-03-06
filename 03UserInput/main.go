package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    welcome:= "welcome to user input"
	fmt.Println(welcome)
	//take input from the users

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating:")

	//comma ok //err err

	input, _ :=reader.ReadString('\n')
	fmt.Println("Thanks for rating",input)
	



}





