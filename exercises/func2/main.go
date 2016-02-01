package main

import "fmt"

func main(){
	divided := func(num int) (int, bool){
		return num/2, num%2==0
	}
	var num int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&num)
	fmt.Println(divided(num))
}
