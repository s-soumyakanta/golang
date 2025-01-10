package main

import "fmt"

func main() {
	fmt.Println("Practice One")
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
	const e, f, g = true, false, "no!"
	fmt.Println(e, f, g)

	const (
		age     = 22
		name    = "Soumya"
		country = "India"
	)

	if age >= 18 || country == "India" {
		fmt.Println(name, "is eligible to vote")
	} else {
		fmt.Println(name, "is not eligible to vote")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i + 10)
	}
}
