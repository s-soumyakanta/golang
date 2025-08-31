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
	fmt.Println(PI, AGE, STATE)
}
