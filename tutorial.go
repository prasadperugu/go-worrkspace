package main
// varible declaration
// var x int & var x,y int
// var name int
// var = keyword
// x = name
// int = type
//short variable declration(works inside a function only)
// x := 10

// week 2: from coursera
// pointers: An address to data in memory
// 1. & operator returns an address of a variable/function
// 2. * operator returns the data at an address(deferencing)
// x:=1;
// var y;
// var ip *int // ip is a pointer to int
// ip = &x // ip now points to x
// y = *ip; // y is now 1


import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type AgifyResponse struct {
	Age int `json:"age"`
}

func main() {
	fmt.Println("hey I havd back")
	var x = 10;
	var y *int;
	y=&x;
	var z = *y;

	
	
	fmt.Println(x);
	fmt.Println(z);
	PredictAgeWithName()

}

func PredictAgeWithName() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Name")
	name, _ := reader.ReadString('\n')
	// formatting string
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	url = strings.Replace(url, "\n", "", -1)
	// sending http request
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	// closing response body
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var agify AgifyResponse
	json.Unmarshal(body, &agify)
	data := []byte(name)
	encodedString := base64.StdEncoding.EncodeToString(data)
	obj := make(map[string]interface{})
	obj["name"] = name
	obj["age"] = agify.Age
	obj["base64"] = encodedString

	fmt.Println("User Info:\n ", obj);
}
