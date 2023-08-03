package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Decelairing variables.
	var a = "variables:"
	fmt.Println(a)

	var num1, num2, num3 int = 1, 2, 3
	fmt.Println(reflect.TypeOf(num1), "Initialized variables:", num1, num2, num3)

	num1, num2, num3 = 11, 24, 35
	fmt.Println(reflect.TypeOf(num1), "Edited variables:", num1, num2, num3)

	var boolTry = true
	fmt.Println(reflect.TypeOf(boolTry), "Initialized variables:", boolTry)

	boolTry = false
	fmt.Println(reflect.TypeOf(boolTry), "Edited variables:", boolTry)

	fruit := "Mango"
	fmt.Println(reflect.TypeOf(fruit), "Initialized variables:", fruit)

	fruit = "Banana"
	fmt.Println(reflect.TypeOf(fruit), "Edited variables:", fruit)

	// Decailing contants.
	const b = "\n\ncontants:"
	fmt.Println(b)

	const test = 54
	fmt.Println(reflect.TypeOf(test), "Initialized constant:", test)

	const testFloat = 54.54
	fmt.Println(reflect.TypeOf(testFloat), "Initialized constant:", testFloat)

	const testStr = "mobile phones"
	fmt.Println(reflect.TypeOf(testStr), "Initialized constant:", testStr)

	// cannot edit constants but we can edit variables.
}
