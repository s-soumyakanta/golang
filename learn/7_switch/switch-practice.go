package main

import "fmt"

func main() {
	x := 2

	switch x {
	case 1:
		println("x is 1")
	case 2:
		println("x is 2")
	default:
		println("x is not 1 or 2")
	}

	var day int = 9

	switch day {
	case 1:
		fmt.Println("Mon")
	case 2:
		fmt.Println("Tue")
	case 3:
		fmt.Println("Wed")
	case 4:
		fmt.Println("Thu")
	case 5:
		fmt.Println("Fri")
	case 6:
		fmt.Println("Sat")
	case 7:
		fmt.Println("Sun")
	default:
		fmt.Println("Invalid Day Number")
	}

	const num = 2
	const even = true

	switch num {
	case 2:
		fmt.Println("true")
	default:
		fmt.Println("not defined")
	}
}
