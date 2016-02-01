package main

import "fmt"

/*asks for a large and small number then print out the remainder of BigNUm/smallNum*/
func main(){

	var num1 int
	var num2 int
	fmt.Print("Enter a large number:")
	fmt.Scan(&num1)
	fmt.Print("Enter a smaller number:")
	fmt. Scan(&num2)
	result := num1%num2
	fmt.Println("The remainder is:", result)
}
