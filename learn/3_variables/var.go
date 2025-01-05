package main

import "fmt"

// We have to use all declaired variables in golang

func main() {
	var name string = "S Soumykanta"
	var langName = "golang"
	var isCool bool = true
	var year int = 2024

	//shorthand syntax - used when to declare variable and assign value to it
	myName := "Soumya"

	//If we are only defining the var then we need to specify its type
	/* If we are declaring a variable in which the value will be assigned
	later then we have to declare the type of the variable at the time of initiation */

	var myNum int
	myNum = 22

	fmt.Println(name, langName, isCool, year, myName, myNum)
}
