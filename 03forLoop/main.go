package main

import "fmt"

func main() {

	// normal for loop.
	fmt.Println("Normal For Loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for loop in range.
	fmt.Println("\nin range of For Loop")
	for k, v := range []string{"a", "b", "c", "d", "e", "f"} {
		fmt.Println("key:", k, "value:", v)
	}

}
