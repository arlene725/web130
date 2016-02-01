package main

/*variadic parameter finds the greatest number in a list of
numbers*/
import "fmt"

func max(numbers ...int)int {

	var largest int
	for _, v := range numbers{
		if v >largest{
			largest = v
		}
	}
	return largest
}


func main(){
	greatest := max(1,2,3,255,6,130,20,30)
	fmt.Println(greatest)
}
