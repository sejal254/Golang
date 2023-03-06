package main

import "fmt"

func main() {

	fmt.Println("Welocme to Struct")

	sejal :=User{"Sejal","sejal@gmail.com",true,21}
	fmt.Println(sejal)
	fmt.Printf("sejal details are: %+v\n",sejal)
	fmt.Printf("Name is %v and email is %v \n.",sejal.Name,sejal.Email)
}


type User struct{
	Name string
	Email string
	Status bool
	Age int

}

func (u User) GetStatus(){
	fmt.Println("Is user active: ",u.Status)

}