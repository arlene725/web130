package main

import "fmt"

/* printing out number 1 to 100, BUT multiples of 3 print "fizz",
multiples of 5 print "buzz", multiples of 3 and 5 printf "fizzbuzz"
 */

func main(){

	for i:= 1; i<= 100; i++{
		if i%3==0 && i%5==0{
			fmt.Println("FizzBuzz")
		}else if i%3 ==0{
			fmt.Println("Fizz")
		}else if  i%5==0{
			fmt.Println("Buzz")
		}else {
			fmt.Println(i)
		}
	}
}