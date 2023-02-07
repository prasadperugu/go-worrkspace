package main
import(
	"fmt"
)
func main(){
	grades := [3]int {10,14,36}
	grades1 := [...] int {35,66,778,77};
	grades1[2] = 55;
	var students [3]string;
	students[1] = "prasad perugu";
	fmt.Println(students[1])

	//another way to define an array without specifying exact length
	//these elements are continuous in memory
	fmt.Println("Grades: %v", grades);
	fmt.Println("Grades: %v", grades1);
	//to get an length of an array
	fmt.Println("length of an array is",len(students));

	a := [...]int {1,4,6,7,78};
	b :=a;
	//will change the original address using that & oeprator
	//b :=&a 
	b[1] = 5;
	fmt.Println(a);
	fmt.Println(b)



	//slices exampe
	//  p := [...] {2,4,5,64,54};//elminate these dots from here that's it;
	 p := []int {3,4,42,44,66,341,1};
	 q := p[:] //slice all the elements
	 r := 	p[3:] //slice from 4th element to end(copied from 4th element)
	 t := p[:3] //slice first 3 elements(copy first 3 elements) 
	 u := p[3:5] //slice the 4th, 5th and 6th elements.
	 fmt.Println(p);
	 fmt.Println(q);
	 fmt.Println(r);
	 fmt.Println(t);
	 fmt.Println(u);


	 fmt.Println(len(p));
	 fmt.Println(cap(p));

	 xyz := make([]int, 3);
	 fmt.Println(xyz)


}
