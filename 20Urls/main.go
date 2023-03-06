//Handling the URLS and Constructing the URLs

package main

import (
	"fmt"
	"net/url"//net is module not just a library is very big and inside it we have small pieces of code
	//like http method as well as url also their
	
)

//construct the Url
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"//base domain+some protocols+portnumber+sub directory+courses


func main() {
	fmt.Println("Welcome to Handling Urls in Golang")
	fmt.Println(myurl)

	//parsing the url
	result, _ :=url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	//query parameters
	qparams:=result.Query()
	//.Query stores all the paremetersinto a better form
	fmt.Printf("The type of query param are:%T \n",qparams)
	fmt.Println(qparams["coursename"])

	for _, val:=range qparams{
		fmt.Println("param is:",val)
	}

	//create url using chunks

	partsofURL:=&url.URL{
		Scheme :"https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=sejal",


	}

	anotherURL:=partsofURL.String()
	fmt.Println(anotherURL)
	





}