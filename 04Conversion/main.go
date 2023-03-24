package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("Welcome to our App")
	fmt.Println("Rate us from 1 to 5")
	reader:= bufio.NewReader(os.Stdin)
	input,_ :=reader.ReadString('\n')
	fmt.Println("Thanks for rating" ,input)
	
	
	i:=42
	f:=float64(i)
	u:=uint(f)
	fmt.Println(u)
	
	
	



}
