package main

import "fmt"

func divided(n int) (int, bool){
	return n/2, n%2 ==0
}
func main(){
	var num int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&num)
	h, even :=divided(num)
	fmt.Println(h, even)

}