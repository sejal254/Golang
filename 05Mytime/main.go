//time package

package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to time package")

	PresentTime :=time.Now()
	fmt.Println(PresentTime)

	fmt.Println(PresentTime.Format("01-02-2006 15:04:05 Monday"))
}