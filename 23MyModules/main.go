package main

//A module is a collection og Go packages stored in a files tree with a go.mod file at its root.

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to the Module Section")
	greeter()
	r:=mux.NewRouter()
	r.HandlerFunc("/",serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000",r))

}

//Go mod init - use to initiliase the package
//Routing System - Gorilla Mux

//install gorilla mux
//with a correctly configured go toolkit
//go get -u github.com/gorilla/mux
//go get is use to directly fetch the things from web, so the above repositary is directly fetch

func greeter(){
	fmt.Println("Hey there mod users")
}

//with the help of r we will get the request
//w respond back 
func serveHome(w http.ResponseWriter,r * http.Request){
	//this serveHome is design for gorilla mux
	w.Write([]byte("<h1>Welcome to Golang</h1>"))

}



