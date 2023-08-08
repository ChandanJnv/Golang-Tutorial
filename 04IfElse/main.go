package main

import "fmt"

func main() {
	fmt.Println("Exploring if else condition")

	// if else condition.
	num := 13
	if num%2 == 0 {
		fmt.Printf("'%v' is even number", num)
	} else {
		fmt.Printf("'%v' is odd number", num)
	}
}
