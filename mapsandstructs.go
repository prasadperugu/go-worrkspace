package main

import(
	"fmt"
)
//defer, panic and recover
//defer: u can invoke a function, but delay in execution in future point of time

func main(){
 //make new
 //new - only allocates - no initiliase of memory
 //make - allocate and initilaise - non zeroed storage
 //map has two things to carry on 1.string 2.value
 // var score map[string]int //it same like new allocating memory no init(nil error-panic)
  score :=make(map[string]int)
 score["prasad"] =848;
 score["perugu"] =839;
  //access the value
  getRunScore := score["prasad"];

 fmt.Println(score);
 fmt.Println(getRunScore);
  delete(score,"perugu");
  fmt.Println(score)

	
}