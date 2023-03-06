package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Verb Video")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	//make a web request
	//this web request may give you a error or a response so write
	response , err := http.Get(myurl)
	if err!=nil{
		panic(err)
	}

	//when all things are done close this
	defer response.Body.Close() 

	fmt.Println("Status code: ",response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)


    var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount ,_:=responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())


//fmt.Println(content) if we directly use this thing the result will we in byte formate
	//fmt.Println(string(content))

//A Builder is used to efficiently build a string using write methods.
//create a builder

	


}