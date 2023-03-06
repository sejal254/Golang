package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working With Files")
	content:= "This content have to put in a file."

	file ,err :=os.Create("./mylocogofile.txt")

	if err !=nil{
		panic(err)
	}
	length, err:=io.WriteString(file,content)

	if err !=nil{
		panic(err)
	}

	fmt.Println("length is: ",length)
	defer file.Close()
	readFile("./mylcogofile.txt")

}

func readFile(filename string){
	databyte, err:=ioutil.ReadFile(filename)
	if err !=nil{
		panic(err)
	}
	fmt.Println("Tect data inside the file is \n",databyte)

}