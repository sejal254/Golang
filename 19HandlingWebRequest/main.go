//lco Request

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main(){
	fmt.Println("LCO web reruest")

	response, err:=http.Get(url)
	if err!=nil{
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n",response)
	defer response.Body.Close()

	datatype ,err :=ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}

	content:=string(datatype)
	fmt.Println(content)

}