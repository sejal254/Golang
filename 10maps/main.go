package main

import "fmt"


// func main() {
// 	//maps also called hashtabel
// 	fmt.Println("welcome to Maps")
// 	//use make to create slices and maps
// 	languages :=make(map[string]string)
// 	languages["JS"]="javascript"
// 	languages["py"]="python"
// 	languages["ry"]="Ruby"

// 	fmt.Println("List of languages: ",languages)
// 	fmt.Println("Js is for",languages["JS"])
// }

type vertex struct{
	l, h float64
}

var m map[string]vertex

func main(){
	m=make(map[string]vertex)
	m["area"]=vertex{
		23.34, 76.34,
	}
	fmt.Println(m["area"])
}

Output----{23.34,76.34}
