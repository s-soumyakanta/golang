package main

import "fmt"

const age = 22

//shorthands not allowed outside of main
// year:=22

var welcomeMessage string = "Hello World"

//Notes

/*Constants are immutable in Go,
meaning their values cannot be changed after they're declared.
*/

/* Constants cannot be declared using the := syntax. */

func main() {
	const myName string = "S Soumyakanta"
	fmt.Println(myName, age, welcomeMessage)

	//multiple constants - constant grouping
	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println(port, host)
}
