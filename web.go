package main

import
(
	"fmt"
	"net/http"
)

const url = "https://teams.microsoft.com/_#/conversations/48:notes?ctx=chat"

func main(){

	fmt.Println("Web Request",url)

	response, err := http.Get(url)

	if err!=nil{
		panic(err)

	}
	fmt.Printf("%T/n",response);
	response.Body.Close()




}