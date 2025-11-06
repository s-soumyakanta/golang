package main

import "fmt"

//while defining a function we need to define intake value type  and it's return value type
// func addNums(a int, b int) int {
//if both parameters are same type then we can define type once like this
func addNums(a, b int) int {
	return a + b
}

//In go functions can return multiple values
//We have to define the types of the return value inside a parenthesis after defining the function
//We generaly return function value and error
func getLanguages() (string, string, int, bool) {
	return "Golang", "TypeScript", 8, true
}

// In go the functions are first class citizens so -
// -> We can assign a function to a variable or
// -> We can pass function as a argument to an anather function

func processIt() func(a int) int {
	return func(a int) int {
		return 9 * a
	}
}

func multiply() func(a, b int) int {
	return func(a, b int) int {
		return a * b
	}
}
func main() {

	fn := processIt()
	fmt.Println(fn(4)) //36
	result := addNums(3, 5)
	//we use _ for not used values
	languages1, languages2, _, _ := getLanguages()
	fmt.Println(languages2, languages1) //TypeScript Golang
	fmt.Println(getLanguages())         //Golang TypeScript 8 true
	fmt.Println(result)                 //8

	mulFn := multiply()
	fmt.Println(mulFn(2, 3)) //6
}
