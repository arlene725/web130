package main

import "fmt"


func main(){
	foo(1,2)
	foo(1,2,3)
	aSlice := []int{1,2,3,4}
	foo(aSlice...)
	foo()
}

/*takes in a list of numbers and prints them out*/

func foo(numbers ...int){
	fmt.Println(numbers)
}