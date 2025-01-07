package main

import "fmt"

const (
	name   = "S Soumyakanta"
	mobile = 2233993377
)

const (
	serverName = "ec2One"
)

const brandName = "Holla"

//shorthands not allowed outside
// dd := 33
const age int = 22

func main() {
	fmt.Println(name, mobile, age)
	fmt.Println(brandName)
	fmt.Println(serverName)
}
