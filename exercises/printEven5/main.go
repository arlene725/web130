package main

/*printing even numbers from 0 to 100*/

import "fmt"

func main(){
	for i:= 0; i<= 100; i++{
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
