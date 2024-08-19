package main

import "fmt"

//Declaring a vaiadic function
//For n numbers of parameters we used 3dots. It similar to JavaScript spread and rest operators
// for any type we use interface{}
// func sum(nums ...interface{}) int {
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}

	return total
}
func main() {
	//Println is a variadic function in go where we can pass n number of arguments
	fmt.Println(1, 2, 3, 4) //1, 2, 3, 4

	result := sum(3, 5, 9)
	fmt.Println(result) //17

	//This is how we can pass values directly from a slice
	someNums := []int{4, 9, 3}
	snResult := sum(someNums...)
	fmt.Println(snResult) //16
}
