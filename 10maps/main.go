package main

import "fmt"

func main() {
	//maps also called hashtabel
	fmt.Println("welcome to Maps")
	//use make to create slices and maps
	languages :=make(map[string]string)
	languages["JS"]="javascript"
	languages["py"]="python"
	languages["ry"]="Ruby"

	fmt.Println("List of languages: ",languages)
	fmt.Println("Js is for",languages["JS"])
}

