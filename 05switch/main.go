package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Exploring switch case condition")

	// switch case condition.
	monthNum := 6
	switch monthNum {
	case 1:
		fmt.Println(time.January)
	case 2:
		fmt.Println(time.February)
	case 3:
		fmt.Println(time.March)
	case 4:
		fmt.Println(time.April)
	case 5:
		fmt.Println(time.May)
	case 6:
		fmt.Println(time.June)
	case 7:
		fmt.Println(time.July)
	case 8:
		fmt.Println(time.August)
	case 9:
		fmt.Println(time.September)
	case 10:
		fmt.Println(time.October)
	case 11:
		fmt.Println(time.November)
	case 12:
		fmt.Println(time.December)
	default:
		fmt.Println("invalid month number")
	}
}
