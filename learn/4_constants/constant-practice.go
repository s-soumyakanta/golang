// package main

// import "fmt"

// const (
// 	name   = "S Soumyakanta"
// 	mobile = 2233993377
// )

// const (
// 	serverName = "ec2One"
// )

// const brandName = "Holla"

// //shorthands not allowed outside
// // dd := 33
// const age int = 22

// func main() {
// 	fmt.Println(name, mobile, age)
// 	fmt.Println(brandName)
// 	fmt.Println(serverName)
// }

package main

import "fmt"

const PI float32 = 3.14

func main() {
	const (
		AGE   int = 20
		STATE     = "Odisha"
	)
	fmt.Printf("AGE has value %v and type: %t", AGE, STATE)
	//Always match the verb (%d, %s, %f, %T, etc.) with the type of the variable.
	fmt.Printf("AGE has value %v and type: %T\n", AGE, AGE)
	fmt.Printf("STATE has value %v and type: %T\n", STATE, STATE)
	fmt.Printf("PI has value %v and type: %T\n", PI, PI)

	fmt.Println(PI, AGE, "\n", STATE)
}
