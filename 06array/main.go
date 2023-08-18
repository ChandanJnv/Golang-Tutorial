package main

import "fmt"

func main() {
	fmt.Println("Exploring arr1ay")

	// initialized array using var.
	var arr1 [5]int
	arr1[0] = 23
	arr1[1] = 546
	arr1[2] = 1324
	arr1[3] = 5436
	arr1[4] = 342
	fmt.Println("array1:", arr1)

	// initialized array directly.
	fruit := [4]string{"mango", "apple", "banana", "date"}
	fmt.Println("fruit:", fruit)

	// 2D array.
	var twoD [2][2]int
	twoD[0][0] = 1
	twoD[0][1] = 2
	twoD[1][0] = 3
	twoD[1][1] = 4
	fmt.Println("2D array:", twoD)

	// 3D array.
	threeD := [3][3]int{{23, 454, 656}, {323, 43, 65}, {12, 34, 4564}}
	fmt.Println("3D array:", threeD)

}
