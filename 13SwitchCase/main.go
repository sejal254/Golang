package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to switch case")
	//making a dice
//using rand to genertae random numbers all the time coming from math/rand

    rand.Seed(time.Now().UnixNano())

//generating random numbers from 1 to 6 as 1 is starting number it is not included so add 1for that
	dicenumber:=rand.Intn(6)+1
	fmt.Println("Value of dice is:",dicenumber)

	switch dicenumber{
	case 1:
		fmt.Println("Dice vale is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	}

	
}