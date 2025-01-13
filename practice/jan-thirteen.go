package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	a := 10

	if a > 11 {
		fmt.Println("a is greater than 11")
	} else if a < 11 {
		fmt.Println("a is less than 11")
	}

	name := "Soumya"

	switch name {
	case "soumya":
		println("Name is soumya")

	case "Soumya":
		println("Name is Soumya")
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	for l := 10; l > 0; l-- {
		fmt.Println(l)
	}

	var (
		age  = 10
		male = true
	)
	fmt.Println(age, male)
	const (
		rollNumber = 44
		section    = "B"
	)
	fmt.Println(rollNumber, section)
}
