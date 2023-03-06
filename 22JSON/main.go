package main

import (
	"encoding/json"
	"fmt"
)

//here we learn about encoding of JSON

type course struct{
	Name string 
	Price int
	Plateform string
	Password string
	Tags []string
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodJSON()
}


func EncodJSON(){
	lcoCourses:=[]course{
		{"React Bootcamp",299,"LearnCodeOnline.in","abc123",[]string{"web-dev","js"}},
		{"MERN Bootcamp",299,"LearnCodeOnline.in","bcd123",[]string{"web-dev","js"}},
		{"Angular Bootcamp",299,"LearnCodeOnline.in","efg123",nil},
	}

	//package this data as JSON data
	finalJSON , err:=json.MarshalIndent(lcoCourses,"","\t")
	if err !=nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJSON)



}

func DecodeJson(){







	
}