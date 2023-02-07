package main

import
(
	"fmt"
	"net/http"
	
)



func main(){

	fmt.Println("Welcome to web application")
	performGetRequest()


}

func performGetRequest(){
	const myurl = "http://localhost:8080"
	 res,err := http.GET(myurl)
	 if err !=nil {
		panic(err)
	 }

	 defer response.Body.Close()
	 fmt.Println("status code:", response.StatusCode)
	 content, _ :=ioutil.ReadAll(response.Body)
	 fmt.Println(string(content))

}

