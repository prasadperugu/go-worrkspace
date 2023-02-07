package main


import(
	"fmt"
	"encoding/json"
)
        //telusko(struct)
		//assign multiple values to one element(varible) is called struct
		//convert struct into json string 
type Student struct {
	roolno int `json:"roolno"`
	name string `json:"name" `
	marks int
}		
func main(){

	var s1 = Student{101,"prasad",70};
	fmt.Println(s1);
	fmt.Println(s1.name);
	var s2 Student = Student{name:"prasad", marks:70};
	fmt.Println(s2);
	byteArray,err := json.Marshal(s2);
	if(err !=nil){
		fmt.Println(err);
	}
	fmt.Println(string(byteArray));




	// fmt.Println("this is me prasad");
	// var arr [5]int;
	// arr[0] = 10;
	// fmt.Println(arr[0]);
	// // var arr1 [5]int = [5] {10,5,677,455,5};
	// // fmt.Println(arr1[3]);
	// // x:=[] int {1,2,3,5};
	// // fmt.Println(x[1]);

	// //slices is a datatype use instead of array becuase u can change their size
	// 	// array := [] string {"a","b","c","d","e","f","g"};
	// 	// s1 := array[1:3]//slice 1(3-1=2)
	// 	// s2 := array[2:5]//slice 2(5-1=4)
	// 	//structs
	// 	//example:person (name,address,phone)=make a single struct which contain all 3 vars
	// 	type struct Person {
	// 		name string;
	// 		addr string;
	// 		phone string;
	// 	}
	// 	var p1 Person;
	// 	//p1 contains values for fileds(name,add,phone)
	// 	//accessing struct fields
	// 		p1.name = "prasad";
	// 		y = p1.addr;
	// 		fmt.Println(y);

	// 		//initializing structs
	// 		// can use new() p1 := new (Person);
	// 		// p1:= Person(name:"prasad",addr:"kovur",phone:"939")
		



}

