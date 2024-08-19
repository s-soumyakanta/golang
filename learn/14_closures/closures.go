package main

import "fmt"

func counter() func() int {
	var count int = 0
	//returning a anonymous function
	return func() int {
		count += 1
		return count
	}
}
func main() {
	increment := counter()
	fmt.Println(increment()) //1
	fmt.Println(increment()) //2
	//This is happening due to closure. Similar to JavaScript
	//It is able find refference of count from its outer scope
}
