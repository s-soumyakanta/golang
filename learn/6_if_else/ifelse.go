package main

import "fmt"

func main() {

	//if - else
	age := 24

	if age >= 18 {
		fmt.Println("Person is an adult!")
	} else {
		fmt.Println("Kid")
	}

	//if - else if
	barAge := 18
	personAge := 18

	if personAge > barAge {
		fmt.Println("Welcome to bar")
	} else if personAge == barAge {
		fmt.Println("You are 18, welcome to bar")
	} else {
		fmt.Println("Not allowed to bar")
	}

	//logical operators within if else
	var role = "admin"
	var hasPermissions = false

	if role == "admin" || hasPermissions {
		fmt.Println("allowed")
	} else {
		fmt.Println("Not allowed")
	}

	// declaring variables in if else
	if age := 14; age >= 18 {
		fmt.Println("Person is adult", age)
	} else if age >= 12 {
		fmt.Println("Person is teenager", age)
	}

	//go does not have ternary operater, we have to use if else

}
