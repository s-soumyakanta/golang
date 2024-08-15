package main

import "fmt"

func main() {
	var name string = "S Soumykanta"
	var langName = "golang"
	var isCool bool = true
	var year int = 2024

	//shorthand syntax - used when to declare variable and assign value to it
	myName := "Soumya"

	//If we are only defining the var then we need to specify its type
	var myNum int
	myNum = 22

	fmt.Println(name, langName, isCool, year, myName, myNum)
}
