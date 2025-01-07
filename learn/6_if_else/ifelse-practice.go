package main

import "fmt"

func main() {
	age := 18

	if age > 18 {
		fmt.Println("Allowed to bar")
	} else if age == 18 {
		fmt.Println("Person age is 18! allowed")
	} else {
		fmt.Println("Not allowed to bar")
	}

	var gender = "female"
	personAge := 3

	if personAge >= 18 || gender == "male" {
		fmt.Println("Welcome to male bar")
	}
}
